#!/bin/bash

set -e

cd "${0%/*}" # cd in the script dir

PEGGY_ID="${PEGGY_ID:-0x666f6f0000000000000000000000000000000000000000000000000000000000}"
POWER_THRESHOLD="${POWER_THRESHOLD:-3500}"
VALIDATOR_ADDRESSES="${VALIDATOR_ADDRESSES:-0xB6362cFC23C88d510AC259FCeDd60a3a34DBf433,0xbAb1da38f751693A4039e101A5DcAdC2ed2e55Ad,0x2c8f0B54Cd3d30B0D38d9978Bc8153d3dB636567}"
VALIDATOR_POWERS="${VALIDATOR_POWERS:-3000,1500,500}"

if [[ ! -f .env ]]; then
	echo "Please create .env file, example is in .env.example"
	exit 1
fi

peggy_impl_address=`evm-deploy-contract \
	--name Peggy \
	--source ../contracts/Peggy.sol \
	deploy`

echo "Deployed Peggy implementation contract: $peggy_impl_address"
echo -e "===\n"

peggy_init_data=`evm-deploy-contract \
	--name Peggy \
	--source ../contracts/Peggy.sol \
	tx --bytecode $peggy_impl_address initialize \
	$PEGGY_ID \
	$POWER_THRESHOLD \
	$VALIDATOR_ADDRESSES \
	$VALIDATOR_POWERS`

echo "Using PEGGY_ID $PEGGY_ID"
echo "Using POWER_THRESHOLD $POWER_THRESHOLD"
echo "Using VALIDATOR_ADDRESSES $VALIDATOR_ADDRESSES"
echo "Using VALIDATOR_POWERS $VALIDATOR_POWERS"
echo -e "===\n"
echo "Peggy Init data: $peggy_init_data"
echo -e "===\n"

proxy_admin_address=`evm-deploy-contract \
	--name ProxyAdmin \
	--source ../contracts/@openzeppelin/contracts/ProxyAdmin.sol \
	deploy`

echo "Deployed ProxyAdmin contract: $proxy_admin_address"
echo -e "===\n"

peggy_proxy_address=`evm-deploy-contract \
	--name TransparentUpgradeableProxy \
	--source ../contracts/@openzeppelin/contracts/TransparentUpgradeableProxy.sol \
	deploy $peggy_impl_address $proxy_admin_address $peggy_init_data`

echo "Deployed TransparentUpgradeableProxy for $peggy_impl_address (Peggy), with $proxy_admin_address (ProxyAdmin) as the admin"
echo -e "===\n"

echo "Peggy deployment done! Use $peggy_proxy_address"
