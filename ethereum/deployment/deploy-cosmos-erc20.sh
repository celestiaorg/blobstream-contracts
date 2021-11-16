#!/bin/bash

set -e

cd "${0%/*}" # cd in the script dir

PEGGY_ADDRESS="${PEGGY_ADDRESS}"

if [[ ! -f .env ]]; then
	echo "Please create .env file, example is in .env.example"
	exit 1
fi

if [[ "$PEGGY_ADDRESS" == "" ]]; then
	echo "Please set \$PEGGY_ADDRESS variable to a deployed Peggy instance"
	exit 1
fi

deploy_erc20_txhash=`etherman \
	--name Peggy \
	--source ../contracts/Peggy.sol \
	tx $PEGGY_ADDRESS deployERC20 inj INJ INJ 18`

deploy_erc20_log=`etherman \
	--name Peggy \
	--source ../contracts/Peggy.sol \
	logs $PEGGY_ADDRESS $deploy_erc20_txhash ERC20DeployedEvent`

erc20_token_address=`jq -r '..|._tokenContract?' <<< $deploy_erc20_log`

echo "Deployed Cosmos ERC20 INJ Contract $erc20_token_address"
