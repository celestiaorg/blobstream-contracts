package cosmos

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/celestiaorg/quantum-gravity-bridge/cmd/peggo/client"
	wrappers "github.com/celestiaorg/quantum-gravity-bridge/ethereum/solidity/wrappers/QuantumGravityBridge.sol"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/keystore"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/peggy"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/umee-network/umee/x/peggy/types"
)

type PeggyBroadcastClient interface {
	ValFromAddress() sdk.ValAddress
	AccFromAddress() sdk.AccAddress

	// Send a transaction updating the eth address for the sending
	// Cosmos address. The sending Cosmos address should be a validator
	UpdatePeggyOrchestratorAddresses(
		ctx context.Context,
		ethFrom ethcmn.Address,
		orchAddr sdk.AccAddress,
	) error

	// SendValsetConfirm broadcasts in a confirmation for a specific validator set for a specific block height.
	SendValsetConfirm(
		ctx context.Context,
		ethFrom ethcmn.Address,
		peggyID ethcmn.Hash,
		valset *types.Valset,
	) error

	// SendBatchConfirm broadcasts in a confirmation for a specific transaction batch set for a specific block height
	// since transaction batches also include validator sets this has all the arguments
	SendBatchConfirm(
		ctx context.Context,
		ethFrom ethcmn.Address,
		peggyID ethcmn.Hash,
		batch *types.OutgoingTxBatch,
	) error

	// SendRequestBatch broadcasts a requests a batch of withdrawal transactions to be generated on the chain.
	SendRequestBatch(
		ctx context.Context,
		denom string,
	) error
}

// sortableEvent exists with the only purpose to make a nicer sortable slice for Ethereum events.
// It is only used in `SendEthereumClaims`
type sortableEvent struct {
	EventNonce        uint64
	TupleRootEvent    *wrappers.QuantumGravityBridgeMessageTupleRootEvent
	ValsetUpdateEvent *wrappers.QuantumGravityBridgeValidatorSetUpdatedEvent
}

func NewPeggyBroadcastClient(
	logger zerolog.Logger,
	queryClient types.QueryClient,
	broadcastClient client.CosmosClient,
	ethSignerFn keystore.SignerFn,
	ethPersonalSignFn keystore.PersonalSignFn,
) PeggyBroadcastClient {
	return &peggyBroadcastClient{
		logger:            logger.With().Str("module", "peggy_broadcast_client").Logger(),
		daemonQueryClient: queryClient,
		broadcastClient:   broadcastClient,
		ethSignerFn:       ethSignerFn,
		ethPersonalSignFn: ethPersonalSignFn,
	}
}

func (s *peggyBroadcastClient) ValFromAddress() sdk.ValAddress {
	return sdk.ValAddress(s.broadcastClient.FromAddress().Bytes())
}

func (s *peggyBroadcastClient) AccFromAddress() sdk.AccAddress {
	return s.broadcastClient.FromAddress()
}

type peggyBroadcastClient struct {
	logger            zerolog.Logger
	daemonQueryClient types.QueryClient
	broadcastClient   client.CosmosClient
	ethSignerFn       keystore.SignerFn
	ethPersonalSignFn keystore.PersonalSignFn
}

func (s *peggyBroadcastClient) UpdatePeggyOrchestratorAddresses(
	ctx context.Context,
	ethFrom ethcmn.Address,
	orchestratorAddr sdk.AccAddress,
) error {
	// SetOrchestratorAddresses

	// This message allows validators to delegate their voting responsibilities
	// to a given key. This key is then used as an optional authentication method
	// for sigining oracle claims
	// This is used by the validators to set the Ethereum address that represents
	// them on the Ethereum side of the bridge. They must sign their Cosmos address
	// using the Ethereum address they have submitted. Like ValsetResponse this
	// message can in theory be submitted by anyone, but only the current validator
	// sets submissions carry any weight.

	// -------------
	msg := &types.MsgSetOrchestratorAddresses{
		Sender:       s.AccFromAddress().String(),
		EthAddress:   ethFrom.Hex(),
		Orchestrator: orchestratorAddr.String(),
	}

	res, err := s.broadcastClient.SyncBroadcastMsg(msg)
	fmt.Fprintf(os.Stderr, "Broadcast MsgSetOrchestratorAddresses response: \n%v\n", res)
	if err != nil {
		err = errors.Wrap(err, "broadcasting MsgSetOrchestratorAddresses failed")
		return err
	}

	return nil
}

func (s *peggyBroadcastClient) SendValsetConfirm(
	ctx context.Context,
	ethFrom ethcmn.Address,
	peggyID ethcmn.Hash,
	valset *types.Valset,
) error {

	confirmHash := peggy.EncodeValsetConfirm(peggyID, valset)
	signature, err := s.ethPersonalSignFn(ethFrom, confirmHash.Bytes())
	if err != nil {
		err = errors.New("failed to sign validator address")
		return err
	}
	// MsgValsetConfirm
	// this is the message sent by the validators when they wish to submit their
	// signatures over the validator set at a given block height. A validator must
	// first call MsgSetEthAddress to set their Ethereum address to be used for
	// signing. Then someone (anyone) must make a ValsetRequest the request is
	// essentially a messaging mechanism to determine which block all validators
	// should submit signatures over. Finally validators sign the validator set,
	// powers, and Ethereum addresses of the entire validator set at the height of a
	// ValsetRequest and submit that signature with this message.
	//
	// If a sufficient number of validators (66% of voting power) (A) have set
	// Ethereum addresses and (B) submit ValsetConfirm messages with their
	// signatures it is then possible for anyone to view these signatures in the
	// chain store and submit them to Ethereum to update the validator set
	// -------------
	msg := &types.MsgValsetConfirm{
		Orchestrator: s.AccFromAddress().String(),
		EthAddress:   ethFrom.Hex(),
		Nonce:        valset.Nonce,
		Signature:    ethcmn.Bytes2Hex(signature),
	}
	if err = s.broadcastClient.QueueBroadcastMsg(msg); err != nil {
		err = errors.Wrap(err, "broadcasting MsgValsetConfirm failed")
		return err
	}

	return nil
}

func (s *peggyBroadcastClient) SendBatchConfirm(
	ctx context.Context,
	ethFrom ethcmn.Address,
	peggyID ethcmn.Hash,
	batch *types.OutgoingTxBatch,
) error {

	confirmHash := peggy.EncodeTxBatchConfirm(peggyID, batch)
	signature, err := s.ethPersonalSignFn(ethFrom, confirmHash.Bytes())
	if err != nil {
		err = errors.New("failed to sign validator address")
		return err
	}

	// MsgConfirmBatch
	// When validators observe a MsgRequestBatch they form a batch by ordering
	// transactions currently in the txqueue in order of highest to lowest fee,
	// cutting off when the batch either reaches a hardcoded maximum size (to be
	// decided, probably around 100) or when transactions stop being profitable
	// (TODO determine this without nondeterminism) This message includes the batch
	// as well as an Ethereum signature over this batch by the validator
	// -------------
	msg := &types.MsgConfirmBatch{
		Orchestrator:  s.AccFromAddress().String(),
		Nonce:         batch.BatchNonce,
		Signature:     ethcmn.Bytes2Hex(signature),
		EthSigner:     ethFrom.Hex(),
		TokenContract: batch.TokenContract,
	}
	if err = s.broadcastClient.QueueBroadcastMsg(msg); err != nil {
		err = errors.Wrap(err, "broadcasting MsgConfirmBatch failed")
		return err
	}

	return nil
}

func (s *peggyBroadcastClient) sendTupleRootClaims(tupleRoot *wrappers.QuantumGravityBridgeMessageTupleRootEvent) error {

	s.logger.Info().
		Str("nonce", tupleRoot.BatchNonce.String()).
		Str("token_contract", tupleRoot.Token.Hex()).
		Str("event_nonce", tupleRoot.EventNonce.String()).
		Msg("oracle observed a tuple root event. Sending MsgTupleRootClaim")

	// TupleRootClaim claims that a tuple root submission
	// operations on the bridge contract was executed.
	msg := &types.MsgWithdrawClaim{
		EventNonce:    tupleRoot.EventNonce.Uint64(),
		BatchNonce:    tupleRoot.BatchNonce.Uint64(),
		BlockHeight:   tupleRoot.Raw.BlockNumber,
		TokenContract: tupleRoot.Token.Hex(),
		Orchestrator:  s.AccFromAddress().String(),
	}

	txResponse, err := s.broadcastClient.SyncBroadcastMsg(msg)
	if err != nil {
		s.logger.Err(err).Msg("broadcasting MsgTupleRootClaim failed")
		return err
	}

	s.logger.Info().
		Str("tx_hash", txResponse.TxHash).
		Str("event_nonce", tupleRoot.EventNonce.String()).
		Msg("oracle sent TupleRoot event successfully")

	return nil
}

func (s *peggyBroadcastClient) sendValsetUpdateClaims(valsetUpdate *wrappers.QuantumGravityBridgeValidatorSetUpdatedEvent) error {

	s.logger.Info().
		Str("event_nonce", valsetUpdate.EventNonce.String()).
		Uint64("valset_nonce", valsetUpdate.NewValsetNonce.Uint64()).
		Int("validators", len(valsetUpdate.Validators)).
		Interface("powers", valsetUpdate.Powers).
		Uint64("reward_amount", valsetUpdate.RewardAmount.Uint64()).
		Str("reward_token", valsetUpdate.RewardToken.Hex()).
		Msg("oracle observed a valset update event; sending MsgValsetUpdateClaim")

	members := make([]*types.BridgeValidator, len(valsetUpdate.Validators))
	for i, val := range valsetUpdate.Validators {
		members[i] = &types.BridgeValidator{
			EthereumAddress: val.Hex(),
			Power:           valsetUpdate.Powers[i].Uint64(),
		}
	}

	msg := &types.MsgValsetUpdatedClaim{
		EventNonce:   valsetUpdate.EventNonce.Uint64(),
		ValsetNonce:  valsetUpdate.NewValsetNonce.Uint64(),
		BlockHeight:  valsetUpdate.Raw.BlockNumber,
		RewardAmount: sdk.NewIntFromBigInt(valsetUpdate.RewardAmount),
		RewardToken:  valsetUpdate.RewardToken.Hex(),
		Members:      members,
		Orchestrator: s.AccFromAddress().String(),
	}

	txResponse, err := s.broadcastClient.SyncBroadcastMsg(msg)
	if err != nil {
		s.logger.Err(err).Msg("broadcasting MsgValsetUpdatedClaim failed")
		return err
	}

	s.logger.Info().
		Str("tx_hash", txResponse.TxHash).
		Str("event_nonce", valsetUpdate.EventNonce.String()).
		Msg("oracle sent ValsetUpdate event successfully")

	return nil
}

func (s *peggyBroadcastClient) SendEthereumClaims(
	ctx context.Context,
	lastClaimEvent uint64,
	tupleRoots []*wrappers.QuantumGravityBridgeMessageTupleRootEvent,
	valsetUpdates []*wrappers.QuantumGravityBridgeValidatorSetUpdatedEvent,
	cosmosBlockTime time.Duration,
) error {
	allevents := []sortableEvent{}

	// We add all the events to the same list to be sorted.
	// Only events that have a nonce higher than the last claim event will be appended.
	for _, ev := range tupleRoots {
		if ev.Nonce.Uint64() > lastClaimEvent {
			allevents = append(allevents, sortableEvent{
				EventNonce:     ev.Nonce.Uint64(),
				TupleRootEvent: ev,
			})
		}
	}

	for _, ev := range valsetUpdates {
		if ev.Nonce.Uint64() > lastClaimEvent {
			allevents = append(allevents, sortableEvent{
				EventNonce:        ev.Nonce.Uint64(),
				ValsetUpdateEvent: ev,
			})
		}
	}

	// Use SliceStable so we always get the same order
	sort.SliceStable(allevents, func(i, j int) bool {
		return allevents[i].EventNonce < allevents[j].EventNonce
	})

	// iterate through events and send them sequentially
	for _, ev := range allevents {
		// If the event nonce isn't sequential, we break from this loop
		// given that the events are sorted, this should never happen.
		if ev.EventNonce != lastClaimEvent+1 {
			break
		}

		switch {
		case ev.TupleRootEvent != nil:
			err := s.sendTupleRootClaims(ev.TupleRootEvent)
			if err != nil {
				s.logger.Err(err).Msg("broadcasting MsgSendTupleRootClaim failed")
				return err
			}
		case ev.ValsetUpdateEvent != nil:
			err := s.sendValsetUpdateClaims(ev.ValsetUpdateEvent)
			if err != nil {
				s.logger.Err(err).Msg("broadcasting MsgValsetUpdateClaim failed")
				return err
			}
		}

		lastClaimEvent++
		time.Sleep(cosmosBlockTime)
	}

	return nil
}

func (s *peggyBroadcastClient) SendRequestBatch(
	ctx context.Context,
	denom string,
) error {
	// MsgRequestBatch
	// this is a message anyone can send that requests a batch of transactions to
	// send across the bridge be created for whatever block height this message is
	// included in. This acts as a coordination point, the handler for this message
	// looks at the AddToOutgoingPool tx's in the store and generates a batch, also
	// available in the store tied to this message. The validators then grab this
	// batch, sign it, submit the signatures with a MsgConfirmBatch before a relayer
	// can finally submit the batch
	// -------------

	msg := &types.MsgRequestBatch{
		Denom:        denom,
		Orchestrator: s.AccFromAddress().String(),
	}
	if err := s.broadcastClient.QueueBroadcastMsg(msg); err != nil {
		err = errors.Wrap(err, "broadcasting MsgRequestBatch failed")
		return err
	}

	return nil
}
