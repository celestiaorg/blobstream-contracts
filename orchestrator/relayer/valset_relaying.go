package relayer

import (
	"context"

	"github.com/pkg/errors"
	"github.com/umee-network/umee/x/peggy/types"
)

// RelayValsets checks the last validator set on Ethereum, if it's lower than our latest validator
// set then we should package and submit the update as an Ethereum transaction
func (s *peggyRelayer) RelayValsets(ctx context.Context) error {
	// we should determine if we need to relay one
	// to Ethereum for that we will find the latest confirmed valset and compare it to the ethereum chain
	latestValsets, err := s.cosmosQueryClient.LatestValsets(ctx)
	if err != nil {
		err = errors.Wrap(err, "failed to fetch latest valsets from cosmos")
		return err
	}

	var latestCosmosSigs []*types.MsgValsetConfirm
	var latestCosmosConfirmed *types.Valset
	for _, set := range latestValsets {
		sigs, err := s.cosmosQueryClient.AllValsetConfirms(ctx, set.Nonce)
		if err != nil {
			err = errors.Wrapf(err, "failed to get valset confirms at nonce %d", set.Nonce)
			return err
		} else if len(sigs) == 0 {
			continue
		}

		latestCosmosSigs = sigs
		latestCosmosConfirmed = set
		break
	}

	if latestCosmosConfirmed == nil {
		s.logger.Debug().Msg("no confirmed valsets found, nothing to relay")
		return nil
	}

	currentEthValset, err := s.FindLatestValset(ctx)
	if err != nil {
		err = errors.Wrap(err, "couldn't find latest confirmed valset on Ethereum")
		return err
	}

	s.logger.Debug().
		Interface("current_eth_valset", currentEthValset).
		Interface("latest_cosmos_confirmed", latestCosmosConfirmed).
		Msg("found latest valsets")

	if latestCosmosConfirmed.Nonce > currentEthValset.Nonce {
		latestEthereumValsetNonce, err := s.peggyContract.GetValsetNonce(ctx, s.peggyContract.FromAddress())
		if err != nil {
			err = errors.Wrap(err, "failed to get latest Valset nonce")
			return err
		}

		// Check if latestCosmosConfirmed already submitted by other validators in mean time
		if latestCosmosConfirmed.Nonce > latestEthereumValsetNonce.Uint64() {
			s.logger.Info().
				Uint64("latest_cosmos_confirmed_nonce", latestCosmosConfirmed.Nonce).
				Uint64("latest_ethereum_valset_nonce", latestEthereumValsetNonce.Uint64()).
				Msg("detected latest cosmos valset nonce, but latest valset on Ethereum is different. Sending update to Ethereum")

			// Send Valset Update to Ethereum
			txHash, err := s.peggyContract.SendEthValsetUpdate(
				ctx,
				currentEthValset,
				latestCosmosConfirmed,
				latestCosmosSigs,
			)
			if err != nil {
				return err
			}

			s.logger.Info().Str("tx_hash", txHash.Hex()).Msg("sent Ethereum Tx (EthValsetUpdate)")

		}

	}

	return nil
}
