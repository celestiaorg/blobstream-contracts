#!/bin/bash

set -e

CWD=$(pwd)

# These options can be overridden by env
GETH_PORT="${GETH_PORT:-8545}"
GETH_NETWORK_ID="${GETH_NETWORK_ID:-50}"
GETH_ALGO="${GETH_ALGO:-ethash}"
GETH_BLOCK_GAS_LIMIT="${GETH_BLOCK_GAS_LIMIT:-60000000}"
CHAIN_DIR="${CHAIN_DIR:-$CWD/data}"

hdir="$CHAIN_DIR/$GETH_NETWORK_ID"

# killall geth
kill $(cat $hdir.geth.pid) &>/dev/null && rm $hdir.geth.pid || true

sleep 1

cd "${0%/*}" # cd to current script dir

if [[ $GETH_ALGO == "ethash" ]]; then
	geth --datadir $hdir --networkid $GETH_NETWORK_ID --nodiscover \
		--http --http.port $GETH_PORT --http.api personal,eth,net,web3 \
		--miner.threads=1 --etherbase=0xBbDf3283d1Cf510c17B4FfA1b900F444bE4A4A4e \
		--mine --targetgaslimit $GETH_BLOCK_GAS_LIMIT > $hdir.geth.log 2>&1 &
	echo $! > $hdir.geth.pid
elif [[ $GETH_ALGO == "clique" ]]; then
	geth --datadir $hdir --networkid $GETH_NETWORK_ID --nodiscover \
		--http --http.port $GETH_PORT --http.api personal,eth,net,web3 --allow-insecure-unlock \
		--unlock 0xBbDf3283d1Cf510c17B4FfA1b900F444bE4A4A4e --password ./geth/clique_password.txt \
		--mine --targetgaslimit $GETH_BLOCK_GAS_LIMIT > $hdir.geth.log 2>&1 &
	echo $! > $hdir.geth.pid
else
	echo "Unsupported Geth algo: $GETH_ALGO, use ethash or clique"
	exit 1
fi

sleep 1

echo
echo "Logs:"
echo "  tail -f ./data/$GETH_NETWORK_ID.geth.log"
echo 
echo "Command Line Access:"
echo "  geth attach http://localhost:8545"
echo "  geth attach ./data/$GETH_NETWORK_ID/geth.ipc"
echo 
echo "Shutdown:"
echo "  kill \$(cat ./data/$GETH_NETWORK_ID.geth.pid)"
