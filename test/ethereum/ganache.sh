#!/bin/bash

set -e

CWD=$(pwd)

# These options can be overridden by env
GANACHE_NETWORK_ID="${GANACHE_NETWORK_ID:-50}"
GANACHE_PORT="${GANACHE_PORT:-8545}"
CHAIN_DIR="${CHAIN_DIR:-$CWD/data}"

hdir="$CHAIN_DIR/$GANACHE_NETWORK_ID"

# kill $(ps -ef | grep ganache-cli | grep -v grep | awk '{print $2}') &>/dev/null || true
kill $(cat $hdir.ganache.pid) &>/dev/null && rm $hdir.ganache.pid || true

sleep 1

ganache-cli --deterministic --networkId $GANACHE_NETWORK_ID -p $GANACHE_PORT --db $hdir/ganache \
	-m 'concert load couple harbor equip island argue ramp clarify fence smart topic' > $hdir.ganache.log 2>&1 &
echo $! > $hdir.ganache.pid

sleep 1

echo
echo "Logs:"
echo "  tail -f ./data/$GANACHE_NETWORK_ID.ganache.log"
echo 
echo "Command Line Access:"
echo "  geth attach http://localhost:8545"
echo 
echo "Shutdown:"
echo "  kill \$(cat ./data/$GANACHE_NETWORK_ID.ganache.pid)"
