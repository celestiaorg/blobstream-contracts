package coingecko

import (
	"math/big"
	"os"
	"testing"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

var logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel).With().Timestamp().Logger()

func TestFeeThresholdTwoDecimals(t *testing.T) {
	// https://api.coingecko.com/api/v3/simple/token_price/ethereum?contract_addresses=0xe28b3b32b6c345a34ff64674606124dd5aceca30&vs_currencies=usd

	injTokenContract := common.HexToAddress("0xe28b3b32b6c345a34ff64674606124dd5aceca30")
	coingeckoFeed := NewCoingeckoPriceFeed(logger, 100, &Config{})
	currentTokenPrice, _ := coingeckoFeed.QueryUSDPrice(injTokenContract) // "usd":9.35

	minFeeInUSD := float64(23.5) // 23.5 USD to submit batch tx
	minInj := minFeeInUSD / currentTokenPrice
	var DecimalReduction = cosmtypes.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// FeeAccumulated is greater than ExpectedFee
	totalFeeInINJ := cosmtypes.NewInt(int64(minInj) + 1).Mul(DecimalReduction)
	isFeeLimitExceeded := coingeckoFeed.CheckFeeThreshold(injTokenContract, totalFeeInINJ, minFeeInUSD)
	assert.True(t, isFeeLimitExceeded, "FeeAccumulated is less than ExpectedFee")

	// FeeAccumulated is less than ExpectedFee
	totalFeeInINJ = cosmtypes.NewInt(int64(minInj) - 1).Mul(DecimalReduction)
	isFeeLimitExceeded = coingeckoFeed.CheckFeeThreshold(injTokenContract, totalFeeInINJ, minFeeInUSD)
	assert.False(t, isFeeLimitExceeded, "FeeAccumulated is greater than ExpectedFee")
}

func TestFeeThresholdNineDecimals(t *testing.T) {
	// https://api.coingecko.com/api/v3/simple/token_price/ethereum?contract_addresses=0x95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce&vs_currencies=usd
	shibTokenContract := common.HexToAddress("0x95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce")
	coingeckoFeed := NewCoingeckoPriceFeed(logger, 100, &Config{})
	currentTokenPrice, _ := coingeckoFeed.QueryUSDPrice(shibTokenContract) // "usd":0.000008853

	minFeeInUSD := float64(23.5) // 23.5 USD to submit batch tx
	minShib := minFeeInUSD / currentTokenPrice
	var DecimalReduction = cosmtypes.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// FeeAccumulated is greater than ExpectedFee
	totalFeeInSHIB := cosmtypes.NewInt(int64(minShib) + 1).Mul(DecimalReduction)
	isFeeLimitExceeded := coingeckoFeed.CheckFeeThreshold(shibTokenContract, totalFeeInSHIB, minFeeInUSD)
	assert.True(t, isFeeLimitExceeded, "FeeAccumulated is less than ExpectedFee")

	// FeeAccumulated is less than ExpectedFee
	totalFeeInSHIB = cosmtypes.NewInt(int64(minShib) - 1).Mul(DecimalReduction)
	isFeeLimitExceeded = coingeckoFeed.CheckFeeThreshold(shibTokenContract, totalFeeInSHIB, minFeeInUSD)
	assert.False(t, isFeeLimitExceeded, "FeeAccumulated is greater than ExpectedFee")
}
