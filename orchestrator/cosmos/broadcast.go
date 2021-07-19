package cosmos

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/sdk-go/chain/client"
	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"

	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/keystore"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/peggy"
	"github.com/InjectiveLabs/peggo/orchestrator/metrics"

	wrappers "github.com/InjectiveLabs/peggo/solidity/wrappers/Peggy.sol"
)

type PeggyBroadcastClient interface {
	ValFromAddress() sdk.ValAddress
	AccFromAddress() sdk.AccAddress

	/// Send a transaction updating the eth address for the sending
	/// Cosmos address. The sending Cosmos address should be a validator
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

	SendEthereumClaims(
		ctx context.Context,
		lastClaimEvent uint64,
		deposits []*wrappers.PeggySendToCosmosEvent,
		withdraws []*wrappers.PeggyTransactionBatchExecutedEvent,
		valsetUpdates []*wrappers.PeggyValsetUpdatedEvent,
	) error

	// SendToEth broadcasts a Tx that tokens from Cosmos to Ethereum.
	// These tokens will not be sent immediately. Instead, they will require
	// some time to be included in a batch.
	SendToEth(
		ctx context.Context,
		destination ethcmn.Address,
		amount, fee sdk.Coin,
	) error

	// SendRequestBatch broadcasts a requests a batch of withdrawal transactions to be generated on the chain.
	SendRequestBatch(
		ctx context.Context,
		denom string,
	) error
}

func NewPeggyBroadcastClient(
	queryClient types.QueryClient,
	broadcastClient client.CosmosClient,
	ethSignerFn keystore.SignerFn,
	ethPersonalSignFn keystore.PersonalSignFn,
) PeggyBroadcastClient {
	return &peggyBroadcastClient{
		daemonQueryClient: queryClient,
		broadcastClient:   broadcastClient,
		ethSignerFn:       ethSignerFn,
		ethPersonalSignFn: ethPersonalSignFn,

		svcTags: metrics.Tags{
			"svc": "peggy_broadcast",
		},
	}
}

func (s *peggyBroadcastClient) ValFromAddress() sdk.ValAddress {
	return sdk.ValAddress(s.broadcastClient.FromAddress().Bytes())
}

func (s *peggyBroadcastClient) AccFromAddress() sdk.AccAddress {
	return s.broadcastClient.FromAddress()
}

type peggyBroadcastClient struct {
	daemonQueryClient types.QueryClient
	broadcastClient   client.CosmosClient
	ethSignerFn       keystore.SignerFn
	ethPersonalSignFn keystore.PersonalSignFn

	svcTags metrics.Tags
}

func (s *peggyBroadcastClient) UpdatePeggyOrchestratorAddresses(
	ctx context.Context,
	ethFrom ethcmn.Address,
	orchestratorAddr sdk.AccAddress,
) error {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()
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
	fmt.Println("Response of set eth address", "res", res)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
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
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	confirmHash := peggy.EncodeValsetConfirm(peggyID, valset)
	signature, err := s.ethPersonalSignFn(ethFrom, confirmHash.Bytes())
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
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
		metrics.ReportFuncError(s.svcTags)
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
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	confirmHash := peggy.EncodeTxBatchConfirm(peggyID, batch)
	signature, err := s.ethPersonalSignFn(ethFrom, confirmHash.Bytes())
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
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
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "broadcasting MsgConfirmBatch failed")
		return err
	}

	return nil
}

func (s *peggyBroadcastClient) sendDepositClaims(
	ctx context.Context,
	deposit *wrappers.PeggySendToCosmosEvent,
) error {
	// EthereumBridgeDepositClaim
	// When more than 66% of the active validator set has
	// claimed to have seen the deposit enter the ethereum blockchain coins are
	// issued to the Cosmos address in question
	// -------------

	log.WithFields(log.Fields{
		"sender":      deposit.Sender.Hex(),
		"destination": sdk.AccAddress(deposit.Destination[12:32]).String(),
		"amount":      deposit.Amount.String(),
		"event_nonce": deposit.EventNonce.String(),
	}).Infoln("Oracle observed a deposit event. Sending MsgDepositClaim")

	msg := &types.MsgDepositClaim{
		EventNonce:     deposit.EventNonce.Uint64(),
		BlockHeight:    deposit.Raw.BlockNumber,
		TokenContract:  deposit.TokenContract.Hex(),
		Amount:         sdk.NewIntFromBigInt(deposit.Amount),
		EthereumSender: deposit.Sender.Hex(),
		CosmosReceiver: sdk.AccAddress(deposit.Destination[12:32]).String(),
		Orchestrator:   s.broadcastClient.FromAddress().String(),
	}

	if txResponse, err := s.broadcastClient.SyncBroadcastMsg(msg); err != nil {
		metrics.ReportFuncError(s.svcTags)
		log.WithError(err).Errorln("broadcasting MsgDepositClaim failed")
		return err
	} else {
		log.WithFields(log.Fields{
			"event_nonce": deposit.EventNonce.String(),
			"txHash":      txResponse.TxHash,
		}).Infoln("Oracle sent deposit event succesfully")
	}

	return nil
}

func (s *peggyBroadcastClient) sendWithdrawClaims(
	ctx context.Context,
	withdraw *wrappers.PeggyTransactionBatchExecutedEvent,
) error {

	log.WithFields(log.Fields{
		"nonce":          withdraw.BatchNonce.String(),
		"token_contract": withdraw.Token.Hex(),
		"event_nonce":    withdraw.EventNonce.String(),
	}).Infoln("Oracle observed a withdraw batch event. Sending MsgWithdrawClaim")

	// WithdrawClaim claims that a batch of withdrawal
	// operations on the bridge contract was executed.
	msg := &types.MsgWithdrawClaim{
		EventNonce:    withdraw.EventNonce.Uint64(),
		BatchNonce:    withdraw.BatchNonce.Uint64(),
		BlockHeight:   withdraw.Raw.BlockNumber,
		TokenContract: withdraw.Token.Hex(),
		Orchestrator:  s.AccFromAddress().String(),
	}

	if txResponse, err := s.broadcastClient.SyncBroadcastMsg(msg); err != nil {
		metrics.ReportFuncError(s.svcTags)
		log.WithError(err).Errorln("broadcasting MsgWithdrawClaim failed")
		return err
	} else {
		log.WithFields(log.Fields{
			"event_nonce": withdraw.EventNonce.String(),
			"txHash":      txResponse.TxHash,
		}).Infoln("Oracle sent Withdraw event succesfully")
	}

	return nil
}

func (s *peggyBroadcastClient) sendValsetUpdateClaims(
	ctx context.Context,
	valsetUpdate *wrappers.PeggyValsetUpdatedEvent,
) error {

	log.WithFields(log.Fields{
		"EventNonce":   valsetUpdate.EventNonce.Uint64(),
		"ValsetNonce":  valsetUpdate.NewValsetNonce.Uint64(),
		"_validators":  valsetUpdate.Validators,
		"_powers":      valsetUpdate.Powers,
		"rewardAmount": valsetUpdate.RewardAmount,
		"rewardToken":  valsetUpdate.RewardToken.Hex(),
	}).Infoln("Oracle observed a valsetUpdate event. Sending MsgValsetUpdatedClaim")

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

	if txResponse, err := s.broadcastClient.SyncBroadcastMsg(msg); err != nil {
		metrics.ReportFuncError(s.svcTags)
		log.WithError(err).Errorln("broadcasting MsgValsetUpdatedClaim failed")
		return err
	} else {
		log.WithFields(log.Fields{
			"event_nonce": valsetUpdate.EventNonce.String(),
			"txHash":      txResponse.TxHash,
		}).Infoln("Oracle sent ValsetUpdate event succesfully")
	}

	return nil
}

func (s *peggyBroadcastClient) SendEthereumClaims(
	ctx context.Context,
	lastClaimEvent uint64,
	deposits []*wrappers.PeggySendToCosmosEvent,
	withdraws []*wrappers.PeggyTransactionBatchExecutedEvent,
	valsetUpdates []*wrappers.PeggyValsetUpdatedEvent,
) error {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()
	totalClaimEvents := len(deposits) + len(withdraws) + len(valsetUpdates)
	var count, i, j, k int

	// Individual arrays (deposits, withdraws, valsetUpdates) are sorted.
	// Broadcast claim events sequentially starting with eventNonce = lastClaimEvent + 1.
	for count < totalClaimEvents {
		if i < len(deposits) && deposits[i].EventNonce.Uint64() == lastClaimEvent+1 {
			// send deposit
			if err := s.sendDepositClaims(ctx, deposits[i]); err != nil {
				metrics.ReportFuncError(s.svcTags)
				log.WithError(err).Errorln("broadcasting MsgDepositClaim failed")
				return err
			}
			i++
		} else if j < len(withdraws) && withdraws[j].EventNonce.Uint64() == lastClaimEvent+1 {
			// send withdraw claim
			if err := s.sendWithdrawClaims(ctx, withdraws[j]); err != nil {
				metrics.ReportFuncError(s.svcTags)
				log.WithError(err).Errorln("broadcasting MsgWithdrawClaim failed")
				return err
			}
			j++
		} else if k < len(valsetUpdates) && valsetUpdates[k].EventNonce.Uint64() == lastClaimEvent+1 {
			// send valset update claim
			if err := s.sendValsetUpdateClaims(ctx, valsetUpdates[k]); err != nil {
				metrics.ReportFuncError(s.svcTags)
				log.WithError(err).Errorln("broadcasting MsgValsetUpdateClaim failed")
				return err
			}
			k++
		}
		count = count + 1
		lastClaimEvent = lastClaimEvent + 1

		// Considering blockTime=2.8s on Injective chain, Adding Sleep to make sure new event is
		// sent only after previous event is executed successfully.
		// Otherwise it will through `non contiguous event nonce` failing CheckTx.
		time.Sleep(3 * time.Second)
	}
	return nil
}

func (s *peggyBroadcastClient) SendToEth(
	ctx context.Context,
	destination ethcmn.Address,
	amount, fee sdk.Coin,
) error {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	// MsgSendToEth
	// This is the message that a user calls when they want to bridge an asset
	// it will later be removed when it is included in a batch and successfully
	// submitted tokens are removed from the users balance immediately
	// -------------
	// AMOUNT:
	// the coin to send across the bridge, note the restriction that this is a
	// single coin not a set of coins that is normal in other Cosmos messages
	// FEE:
	// the fee paid for the bridge, distinct from the fee paid to the chain to
	// actually send this message in the first place. So a successful send has
	// two layers of fees for the user
	msg := &types.MsgSendToEth{
		Sender:    s.AccFromAddress().String(),
		EthDest:   destination.Hex(),
		Amount:    amount,
		BridgeFee: fee, // TODO: use exactly that fee for transaction
	}
	if err := s.broadcastClient.QueueBroadcastMsg(msg); err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "broadcasting MsgSendToEth failed")
		return err
	}

	return nil
}

func (s *peggyBroadcastClient) SendRequestBatch(
	ctx context.Context,
	denom string,
) error {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

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
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "broadcasting MsgRequestBatch failed")
		return err
	}

	return nil
}
