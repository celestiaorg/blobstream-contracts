package relayer

import (
	"context"

	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/peggo/modules/peggy/types"
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
			err = errors.Wrapf(err, "failed to get valset confims at nonce %d", set.Nonce)
			return err
		} else if len(sigs) == 0 {
			continue
		}

		latestCosmosSigs = sigs
		latestCosmosConfirmed = set
		break
	}

	if latestCosmosConfirmed == nil {
		log.Warningln("no confirmed valsets found, nothing to relay")
		return nil
	}

	currentEthValset, err := s.FindLatestValset(ctx)
	if err != nil {
		err = errors.Wrap(err, "couldn't find latest confirmed valset on Ethereum")
		return err
	}
	log.Debugln("Found Latest valset", "currentEthValset", currentEthValset)

	if latestCosmosConfirmed.Nonce > currentEthValset.Nonce {
		log.Infoln("Detected latest cosmos valset nonce %d, but latest on Ehtereum is %d. Sending update",
			latestCosmosConfirmed.Nonce, currentEthValset.Nonce)

		// TODO(xlab): if the power difference is less than one percent, skip updating
		// the validator set

		txHash, err := s.peggyContract.SendEthValsetUpdate(
			ctx,
			currentEthValset,
			latestCosmosConfirmed,
			latestCosmosSigs,
		)
		if err != nil {
			return err
		}

		log.WithField("tx_hash", txHash.Hex()).Infoln("Sent Ethereum Tx (EthValsetUpdate)")
	}

	return nil
}
