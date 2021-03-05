#!/bin/bash

set -e

CWD=$(pwd)

HARDHAT_NETWORK_ID=50

# These options can be overridden by env
HARDHAT_PORT="${HARDHAT_PORT:-8545}"
CHAIN_DIR="${CHAIN_DIR:-$CWD/data}"

hdir="$CHAIN_DIR/$HARDHAT_NETWORK_ID"

# kill $(ps -ef | grep hardhat | grep -v grep | awk '{print $2}') &>/dev/null || true
kill $(cat $hdir.hardhat.pid) &>/dev/null && rm $hdir.hardhat.pid || true

sleep 1

cd "${0%/*}" # cd to current script dir
yarn hardhat node --port $HARDHAT_PORT > $hdir.hardhat.log 2>&1 &
echo $! > $hdir.hardhat.pid

sleep 1

echo
echo "Logs:"
echo "  tail -f ./data/$HARDHAT_NETWORK_ID.hardhat.log"
echo 
echo "Shutdown:"
echo "  kill \$(cat ./data/$HARDHAT_NETWORK_ID.hardhat.pid)"
