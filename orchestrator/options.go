package orchestrator

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/umee-network/peggo/orchestrator/coingecko"
)

// SetMinBatchFee sets the (optional) minimum batch fee denominated in USD.
func SetMinBatchFee(minFee float64) func(PeggyOrchestrator) {
	return func(p PeggyOrchestrator) { p.SetMinBatchFee(minFee) }
}

func (p *peggyOrchestrator) SetMinBatchFee(minFee float64) {
	p.minBatchFeeUSD = minFee
}

// SetERC20ContractMapping sets the (optional) ERC20 contract mapping Cosmos
// native token denominations to their ERC20 contract addresses.
func SetERC20ContractMapping(m map[ethcmn.Address]string) func(PeggyOrchestrator) {
	return func(p PeggyOrchestrator) { p.SetERC20ContractMapping(m) }
}

func (p *peggyOrchestrator) SetERC20ContractMapping(m map[ethcmn.Address]string) {
	p.erc20ContractMapping = m
}

// SetPriceFeeder sets the (optional) price feeder used when performing profitable
// batch calculations. Note, this should be supplied only when the min batch
// fee is non-zero.
func SetPriceFeeder(pf *coingecko.CoingeckoPriceFeed) func(PeggyOrchestrator) {
	return func(p PeggyOrchestrator) { p.SetPriceFeeder(pf) }
}

func (p *peggyOrchestrator) SetPriceFeeder(pf *coingecko.CoingeckoPriceFeed) {
	p.priceFeeder = pf
}
