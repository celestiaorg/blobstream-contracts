package relayer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

// relayBatches checks the last validator set on Ethereum, if it's lower than our latest valida
// set then we should package and submit the update as an Ethereum transaction
func (s *peggyRelayer) relayBatches(ctx context.Context) error {
	latestBatches, err := s.cosmosQueryClient.LatestTransactionBatches(ctx)
	if err != nil {
		return err
	}

	var oldestSignedBatch *types.OutgoingTxBatch
	var oldestSigs []*types.MsgConfirmBatch
	for _, batch := range latestBatches {
		sigs, err := s.cosmosQueryClient.TransactionBatchSignatures(ctx, batch.BatchNonce, common.HexToAddress(batch.TokenContract))
		if err != nil {
			return err
		} else if len(sigs) == 0 {
			continue
		}

		oldestSignedBatch = batch
		oldestSigs = sigs
	}
	if oldestSignedBatch == nil {
		log.Warningln("could not find batch with signatures, nothing to relay")
		return nil
	}

	latestEthereumBatch, err := s.peggyContract.GetTxBatchNonce(
		ctx,
		common.HexToAddress(oldestSignedBatch.TokenContract),
		s.peggyContract.FromAddress(),
	)
	if err != nil {
		return err
	}

	currentValset, err := s.findLatestValset(ctx)
	if err != nil {
		return errors.New("failed to find latest valset")
	} else if currentValset == nil {
		return errors.New("latest valset not found")
	}

	log.Debugln("Found Latest valset", "currentValset", currentValset)
	if oldestSignedBatch.BatchNonce > latestEthereumBatch.Uint64() {
		log.Infof("We have detected latest batch %d but latest on Ethereum is %d sending an update!", oldestSignedBatch.BatchNonce, latestEthereumBatch)

		txHash, err := s.peggyContract.SendTransactionBatch(ctx, currentValset, oldestSignedBatch, oldestSigs)
		if err != nil {
			return err
		}

		log.WithField("tx_hash", txHash.Hex()).Infoln("Sent Ethereum Tx (TransactionBatch)")
	}

	return nil
}
