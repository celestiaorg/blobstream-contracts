package relayer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/umee-network/umee/x/peggy/types"
)

// RelayBatches checks the last validator set on Ethereum, if it's lower than our latest valida
// set then we should package and submit the update as an Ethereum transaction
func (s *peggyRelayer) RelayBatches(ctx context.Context) error {
	latestBatches, err := s.cosmosQueryClient.LatestTransactionBatches(ctx)
	if err != nil {
		return err
	}

	var oldestSignedBatch *types.OutgoingTxBatch
	var oldestSigs []*types.MsgConfirmBatch
	for _, batch := range latestBatches {
		sigs, err := s.cosmosQueryClient.TransactionBatchSignatures(
			ctx,
			batch.BatchNonce,
			common.HexToAddress(batch.TokenContract),
		)
		if err != nil {
			return err
		} else if len(sigs) == 0 {
			continue
		}

		oldestSignedBatch = batch
		oldestSigs = sigs
	}
	if oldestSignedBatch == nil {
		s.logger.Debug().Msg("could not find batch with signatures, nothing to relay")
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

	currentValset, err := s.FindLatestValset(ctx)
	if err != nil {
		return errors.New("failed to find latest valset")
	} else if currentValset == nil {
		return errors.New("latest valset not found")
	}

	s.logger.Debug().
		Uint64("oldest_batch_nonce", oldestSignedBatch.BatchNonce).
		Uint64("latest_batch_nonce", latestEthereumBatch.Uint64()).
		Msg("found latest valsets")

	if oldestSignedBatch.BatchNonce > latestEthereumBatch.Uint64() {

		latestEthereumBatch, err := s.peggyContract.GetTxBatchNonce(
			ctx,
			common.HexToAddress(oldestSignedBatch.TokenContract),
			s.peggyContract.FromAddress(),
		)
		if err != nil {
			return err
		}
		// Check if oldestSignedBatch already submitted by other validators in mean time
		if oldestSignedBatch.BatchNonce > latestEthereumBatch.Uint64() {
			s.logger.Info().
				Uint64("latest_batch", oldestSignedBatch.BatchNonce).
				Uint64("latest_ethereum_batch", latestEthereumBatch.Uint64()).
				Msg("we have detected latest batch but Ethereum has a different one. Sending an update!")

			// Send SendTransactionBatch to Ethereum
			txHash, err := s.peggyContract.SendTransactionBatch(ctx, currentValset, oldestSignedBatch, oldestSigs)
			if err != nil {
				return err
			}
			s.logger.Info().Str("tx_hash", txHash.Hex()).Msg("sent Ethereum Tx (TransactionBatch)")
		}
	}

	return nil
}
