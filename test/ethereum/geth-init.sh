#!/bin/bash

set -e

CWD=$(pwd)

# These options can be overridden by env
GETH_NETWORK_ID="${GETH_NETWORK_ID:-50}"
GETH_ALGO="${GETH_ALGO:-ethash}"
CHAIN_DIR="${CHAIN_DIR:-$CWD/data}"

hdir="$CHAIN_DIR/$GETH_NETWORK_ID"
ddir="--datadir $hdir"

cd "${0%/*}" # cd to current script dir

if [[ $GETH_ALGO == "ethash" ]]; then
	geth init $ddir ./geth/genesis.json
elif [[ $GETH_ALGO == "clique" ]]; then
	geth init $ddir ./geth/clique_genesis.json
	geth account import $ddir --lightkdf --password ./geth/clique_password.txt ./geth/clique_signer.key
else
	echo "Unsupported Geth algo: $GETH_ALGO, use ethash or clique"
	exit 1
fi
