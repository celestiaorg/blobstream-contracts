package main

import (
	cli "github.com/jawher/mow.cli"
)

var (
	cosmosPrivkey *string
	cosmosGRPC    *string
	tendermintRPC *string
	feeDenom      *string
	ethPrivKey    *string
	chainId       *string
)

func initFlags() {
	cosmosPrivkey = app.String(cli.StringOpt{
		Name:   "cosmos-privkey",
		Desc:   "The Cosmos private key of the validator. Must be saved when you generate your key",
		EnvVar: "PEGGY_COSMOS_PRIVKEY",
	})

	cosmosGRPC = app.String(cli.StringOpt{
		Name:   "cosmos-grpc",
		Desc:   "Cosmos GRPC querying endpoint",
		EnvVar: "PEGGY_COSMOS_GRPC",
		Value:  "tcp://localhost:9900",
	})

	tendermintRPC = app.String(cli.StringOpt{
		Name:   "tendermint-rpc",
		Desc:   "Tednermint RPC endpoint",
		EnvVar: "PEGGY_TENDERMINT_RPC",
		Value:  "http://localhost:26657",
	})

	feeDenom = app.String(cli.StringOpt{
		Name:   "fees",
		Desc:   "The Cosmos Denom in which to pay Cosmos chain fees",
		EnvVar: "PEGGY_FEE_DENOM",
		Value:  "inj",
	})

	ethPrivKey = app.String(cli.StringOpt{
		Name:   "eth-privkey",
		Desc:   "The Ethereum private key of the validator(Ex: 5D862464FE95...)",
		EnvVar: "PEGGY_ETH_PRIVATE_KEY",
	})

	chainId = app.String(cli.StringOpt{
		Name:   "chain-id",
		Desc:   "Specify Chain ID of the injectived service.",
		EnvVar: "INJECTIVED_CHAIN_ID",
		Value:  "888",
	})

}
