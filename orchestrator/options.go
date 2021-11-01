package orchestrator

import (
	"github.com/umee-network/peggo/orchestrator/coingecko"
)

// SetMinBatchFee sets the (optional) minimum batch fee denominated in USD.
func SetMinBatchFee(minFee float64) func(PeggyOrchestrator) {
	return func(p PeggyOrchestrator) { p.SetMinBatchFee(minFee) }
}

func (p *peggyOrchestrator) SetMinBatchFee(minFee float64) {
	p.minBatchFeeUSD = minFee
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
