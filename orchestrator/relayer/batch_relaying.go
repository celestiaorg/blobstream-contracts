package relayer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"
	"github.com/umee-network/peggo/orchestrator/metrics"
)

// RelayBatches checks the last validator set on Ethereum, if it's lower than our latest valida
// set then we should package and submit the update as an Ethereum transaction
func (s *peggyRelayer) RelayBatches(ctx context.Context) error {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	latestBatches, err := s.cosmosQueryClient.LatestTransactionBatches(ctx)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return err
	}
	var oldestSignedBatch *types.OutgoingTxBatch
	var oldestSigs []*types.MsgConfirmBatch
	for _, batch := range latestBatches {
		sigs, err := s.cosmosQueryClient.TransactionBatchSignatures(ctx, batch.BatchNonce, common.HexToAddress(batch.TokenContract))
		if err != nil {
			metrics.ReportFuncError(s.svcTags)
			return err
		} else if len(sigs) == 0 {
			continue
		}

		oldestSignedBatch = batch
		oldestSigs = sigs
	}
	if oldestSignedBatch == nil {
		log.Debugln("could not find batch with signatures, nothing to relay")
		return nil
	}

	latestEthereumBatch, err := s.peggyContract.GetTxBatchNonce(
		ctx,
		common.HexToAddress(oldestSignedBatch.TokenContract),
		s.peggyContract.FromAddress(),
	)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return err
	}

	currentValset, err := s.FindLatestValset(ctx)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return errors.New("failed to find latest valset")
	} else if currentValset == nil {
		metrics.ReportFuncError(s.svcTags)
		return errors.New("latest valset not found")
	}

	log.WithFields(log.Fields{"oldestSignedBatchNonce": oldestSignedBatch.BatchNonce, "latestEthereumBatchNonce": latestEthereumBatch.Uint64()}).Debugln("Found Latest valsets")

	if oldestSignedBatch.BatchNonce > latestEthereumBatch.Uint64() {

		latestEthereumBatch, err := s.peggyContract.GetTxBatchNonce(
			ctx,
			common.HexToAddress(oldestSignedBatch.TokenContract),
			s.peggyContract.FromAddress(),
		)
		if err != nil {
			metrics.ReportFuncError(s.svcTags)
			return err
		}
		// Check if oldestSignedBatch already submitted by other validators in mean time
		if oldestSignedBatch.BatchNonce > latestEthereumBatch.Uint64() {
			log.Infof("We have detected latest batch %d but latest on Ethereum is %d sending an update!", oldestSignedBatch.BatchNonce, latestEthereumBatch)

			// Send SendTransactionBatch to Ethereum
			txHash, err := s.peggyContract.SendTransactionBatch(ctx, currentValset, oldestSignedBatch, oldestSigs)
			if err != nil {
				metrics.ReportFuncError(s.svcTags)
				return err
			}
			log.WithField("tx_hash", txHash.Hex()).Infoln("Sent Ethereum Tx (TransactionBatch)")
		}
	}

	return nil
}
