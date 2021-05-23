package coingecko

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"strings"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func CheckFeeThreshod(erc20Contract common.Address, totalFee cosmtypes.Int, minFeeInUSD float64) bool {
	tokenPriceInUSD := QueryUSDPrice(erc20Contract)
	// Covert token price to Dec with precision of 9 decimals
	tokenPriceWithPrec := int(tokenPriceInUSD * float64(1000000000))
	tokenPriceDecWithPrec := cosmtypes.NewDecWithPrec(int64(tokenPriceWithPrec), 9)
	totalFeeInUSD := cosmtypes.NewDecFromInt(totalFee).Mul(tokenPriceDecWithPrec)

	// Expected minimum fee with precision of 9 decimals
	minFeeInUSDWithPrec := int(minFeeInUSD * float64(1000000000))
	minFeeInUSDDec := cosmtypes.NewDecWithPrec(int64(minFeeInUSDWithPrec), 9)
	var DecimalReduction = cosmtypes.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	minFeeInUSDDecReduced := cosmtypes.NewDecFromInt(DecimalReduction).Mul(minFeeInUSDDec)

	if totalFeeInUSD.GT(minFeeInUSDDecReduced) {
		return true
	}
	return false
}

func QueryUSDPrice(erc20Contract common.Address) float64 {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/token_price/ethereum?contract_addresses=" + strings.ToLower(erc20Contract.String()) + "&vs_currencies=usd")
	defer resp.Body.Close()
	if err != nil {
		return float64(0)
	}
	body, err := io.ReadAll(resp.Body)

	var f interface{}
	err = json.Unmarshal(body, &f)
	m := f.(map[string]interface{})

	v := m[strings.ToLower(erc20Contract.String())]
	n := v.(map[string]interface{})

	tokenPriceInUSD := n["usd"].(float64)
	return tokenPriceInUSD
}
