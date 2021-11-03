package coingecko

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

const (
	maxRespTime        = 15 * time.Second
	maxRespHeadersTime = 15 * time.Second
	maxRespBytes       = 10 * 1024 * 1024
)

var zeroPrice = float64(0)

type PriceFeed struct {
	client *http.Client
	config *Config

	interval time.Duration

	logger zerolog.Logger
}

type Config struct {
	BaseURL string
}

func urlJoin(baseURL string, segments ...string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(append([]string{u.Path}, segments...)...)
	return u.String()

}

func (cp *PriceFeed) QueryUSDPrice(erc20Contract common.Address) (float64, error) {

	u, err := url.ParseRequestURI(urlJoin(cp.config.BaseURL, "simple", "token_price", "ethereum"))
	if err != nil {
		cp.logger.Fatal().Err(err).Msg("failed to parse URL")
	}

	q := make(url.Values)

	q.Set("contract_addresses", strings.ToLower(erc20Contract.String()))
	q.Set("vs_currencies", "usd")
	u.RawQuery = q.Encode()

	reqURL := u.String()
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		cp.logger.Fatal().Err(err).Msg("failed to create HTTP request")
	}

	resp, err := cp.client.Do(req)
	if err != nil {
		err = errors.Wrapf(err, "failed to fetch price from %s", reqURL)
		return zeroPrice, err
	}

	respBody, err := ioutil.ReadAll(io.LimitReader(resp.Body, maxRespBytes))
	if err != nil {
		_ = resp.Body.Close()
		err = errors.Wrapf(err, "failed to read response body from %s", reqURL)
		return zeroPrice, err
	}
	_ = resp.Body.Close()

	var f interface{}
	err = json.Unmarshal(respBody, &f)
	if err != nil {
		return zeroPrice, errors.Wrapf(err, "failed to parse response body from %s", reqURL)
	}

	m := f.(map[string]interface{})

	v := m[strings.ToLower(erc20Contract.String())]
	n := v.(map[string]interface{})

	tokenPriceInUSD := n["usd"].(float64)
	return tokenPriceInUSD, nil
}

// NewCoingeckoPriceFeed returns price puller for given symbol. The price will be pulled
// from endpoint and divided by scaleFactor. Symbol name (if reported by endpoint) must match.
func NewCoingeckoPriceFeed(logger zerolog.Logger, interval time.Duration, endpointConfig *Config) *PriceFeed {
	return &PriceFeed{
		client: &http.Client{
			Transport: &http.Transport{
				ResponseHeaderTimeout: maxRespHeadersTime,
			},
			Timeout: maxRespTime,
		},
		config: checkCoingeckoConfig(endpointConfig),

		interval: interval,

		logger: logger.With().Str("module", "coingecko_pricefeed").Logger(),
	}
}

func checkCoingeckoConfig(cfg *Config) *Config {
	if cfg == nil {
		cfg = &Config{}
	}

	if len(cfg.BaseURL) == 0 {
		cfg.BaseURL = "https://api.coingecko.com/api/v3"
	}

	return cfg
}

func (cp *PriceFeed) CheckFeeThreshold(erc20Contract common.Address, totalFee cosmtypes.Int, minFeeInUSD float64) bool {
	tokenPriceInUSD, err := cp.QueryUSDPrice(erc20Contract)
	if err != nil {
		return false
	}

	tokenPriceInUSDDec := decimal.NewFromFloat(tokenPriceInUSD)
	totalFeeInUSDDec := decimal.NewFromBigInt(totalFee.BigInt(), -18).Mul(tokenPriceInUSDDec)
	minFeeInUSDDec := decimal.NewFromFloat(minFeeInUSD)

	return totalFeeInUSDDec.GreaterThan(minFeeInUSDDec)

}
