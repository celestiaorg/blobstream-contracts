package peggy

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/peggo/orchestrator/metrics"
	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

type ValsetArgs struct {
	Validators   []common.Address `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators,omitempty"`
	Powers       []*big.Int       `protobuf:"varint,1,opt,name=powers,proto3" json:"powers,omitempty"`
	ValsetNonce  *big.Int         `protobuf:"varint,3,opt,name=valsetNonce,proto3" json:"valsetNonce,omitempty"`
	RewardAmount *big.Int         `protobuf:"bytes,4,opt,name=rewardAmount,json=rewardAmount,proto3" json:"rewardAmount"`
	// the reward token in it's Ethereum hex address representation
	RewardToken common.Address `protobuf:"bytes,5,opt,name=rewardToken,json=rewardToken,proto3" json:"rewardToken,omitempty"`
}

func (s *peggyContract) SendEthValsetUpdate(
	ctx context.Context,
	oldValset *types.Valset,
	newValset *types.Valset,
	confirms []*types.MsgValsetConfirm,
) (*common.Hash, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	if newValset.Nonce <= oldValset.Nonce {
		metrics.ReportFuncError(s.svcTags)
		err := errors.New("new valset nonce should be greater than old valset nonce")
		return nil, err
	}

	log.WithFields(log.Fields{
		"old_nonce": oldValset.Nonce,
		"new_nonce": newValset.Nonce,
	}).Infoln("Checking signatures and submitting validator set update to Ethereum")

	newValidators, newPowers := validatorsAndPowers(newValset)
	newValsetNonce := new(big.Int).SetUint64(newValset.Nonce)

	newValsetArgs := ValsetArgs{
		Validators:   newValidators,
		Powers:       newPowers,
		ValsetNonce:  newValsetNonce,
		RewardAmount: newValset.RewardAmount.BigInt(),
		RewardToken:  common.HexToAddress(newValset.RewardToken),
	}

	// we need to use the old valset here because our signatures need to match the current
	// members of the validator set in the contract.
	currentValidators, currentPowers, sigV, sigR, sigS, err := checkValsetSigsAndRepack(oldValset, confirms)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "confirmations check failed")
		return nil, err
	}
	currentValsetNonce := new(big.Int).SetUint64(oldValset.Nonce)
	currentValsetArgs := ValsetArgs{
		Validators:   currentValidators,
		Powers:       currentPowers,
		ValsetNonce:  currentValsetNonce,
		RewardAmount: oldValset.RewardAmount.BigInt(),
		RewardToken:  common.HexToAddress(oldValset.RewardToken),
	}
	// Solidity function signature
	// function updateValset(
	// 		// The new version of the validator set
	// 		address[] memory _newValidators,
	// 		uint256[] memory _newPowers,
	// 		uint256 _newValsetNonce,
	//
	// 		// The current validators that approve the change
	// 		address[] memory _currentValidators,
	// 		uint256[] memory _currentPowers,
	// 		uint256 _currentValsetNonce,
	//
	// 		// These are arrays of the parts of the current validator's signatures
	// 		uint8[] memory _v,
	// 		bytes32[] memory _r,
	// 		bytes32[] memory _s
	// )
	log.Debugln("Sending updateValset Ethereum tx", "currentValidators", currentValidators, "currentPowers", currentPowers, "currentValsetNonce", currentValsetNonce)
	txData, err := peggyABI.Pack("updateValset",
		newValsetArgs,
		currentValsetArgs,
		sigV,
		sigR,
		sigS,
	)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		log.WithError(err).Errorln("ABI Pack (Peggy updateValset) method")
		return nil, err
	}

	// Checking in pending txs(mempool) if tx with same input is already submitted
	if s.pendingTxInputList.IsPendingTxInput(string(txData)) {
		return nil, errors.New("Transaction with same valset input data is already present in mempool")
	}

	txHash, err := s.SendTx(ctx, s.peggyAddress, txData)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		log.WithError(err).WithField("tx_hash", txHash.Hex()).Errorln("Failed to sign and submit (Peggy updateValset) to EVM")
		return nil, err
	}

	log.Infoln("Sent Tx (Peggy updateValset):", txHash.Hex())

	//     let before_nonce = get_valset_nonce(peggy_contract_address, eth_address, web3).await?;
	//     if before_nonce != old_nonce {
	//         info!(
	//             "Someone else updated the valset to {}, exiting early",
	//             before_nonce
	//         );
	//         return Ok(());
	//     }

	//     let tx = web3
	//         .send_transaction(
	//             peggy_contract_address,
	//             payload,
	//             0u32.into(),
	//             eth_address,
	//             our_eth_key,
	//             vec![SendTxOption::GasLimit(1_000_000u32.into())],
	//         )
	//         .await?;
	//     info!("Sent valset update with txid {:#066x}", tx);

	//     // TODO this segment of code works around the race condition for submitting valsets mostly
	//     // by not caring if our own submission reverts and only checking if the valset has been updated
	//     // period not if our update succeeded in particular. This will require some further consideration
	//     // in the future as many independent relayers racing to update the same thing will hopefully
	//     // be the common case.
	//     web3.wait_for_transaction(tx, timeout, None).await?;

	//     let last_nonce = get_valset_nonce(peggy_contract_address, eth_address, web3).await?;
	//     if last_nonce != new_nonce {
	//         error!(
	//             "Current nonce is {} expected to update to nonce {}",
	//             last_nonce, new_nonce
	//         );
	//     } else {
	//         info!(
	//             "Successfully updated Valset with new Nonce {:?}",
	//             last_nonce
	//         );
	//     }
	//     Ok(())

	return &txHash, nil
}

func validatorsAndPowers(valset *types.Valset) (
	validators []common.Address,
	powers []*big.Int,
) {
	for _, m := range valset.Members {
		mPower := big.NewInt(0).SetUint64(m.Power)
		validators = append(validators, common.HexToAddress(m.EthereumAddress))
		powers = append(powers, mPower)
	}

	return
}

func checkValsetSigsAndRepack(
	valset *types.Valset,
	confirms []*types.MsgValsetConfirm,
) (
	validators []common.Address,
	powers []*big.Int,
	v []uint8,
	r []common.Hash,
	s []common.Hash,
	err error,
) {
	if len(confirms) == 0 {
		err = errors.New("no signatures in valset confirmation")
		return
	}

	signerToSig := make(map[string]*types.MsgValsetConfirm, len(confirms))
	for _, sig := range confirms {
		signerToSig[sig.EthAddress] = sig
	}

	powerOfGoodSigs := new(big.Int)
	for _, m := range valset.Members {
		mPower := big.NewInt(0).SetUint64(m.Power)
		if sig, ok := signerToSig[m.EthereumAddress]; ok && sig.EthAddress == m.EthereumAddress {
			powerOfGoodSigs.Add(powerOfGoodSigs, mPower)
			validators = append(validators, common.HexToAddress(m.EthereumAddress))
			powers = append(powers, mPower)
			sigV, sigR, sigS := sigToVRS(sig.Signature)
			v = append(v, sigV)
			r = append(r, sigR)
			s = append(s, sigS)
		} else {
			validators = append(validators, common.HexToAddress(m.EthereumAddress))
			powers = append(powers, mPower)
			v = append(v, 0)
			r = append(r, [32]byte{})
			s = append(s, [32]byte{})
		}
	}
	if peggyPowerToPercent(powerOfGoodSigs) < 66 {
		err = ErrInsufficientVotingPowerToPass
		return
	}

	return
}
