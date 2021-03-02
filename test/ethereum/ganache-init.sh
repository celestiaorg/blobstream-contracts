#!/bin/bash

set -e

CWD=$(pwd)

# These options can be overridden by env
GANACHE_NETWORK_ID="${GANACHE_NETWORK_ID:-50}"
CHAIN_DIR="${CHAIN_DIR:-$CWD/data}"

hdir="$CHAIN_DIR/$GANACHE_NETWORK_ID"
mkdir -p $hdir

wget http://ganache-snapshots.0x.org.s3-website.us-east-2.amazonaws.com/0x_ganache_snapshot-latest.zip && \
	unzip 0x_ganache_snapshot-latest.zip && \
	rm 0x_ganache_snapshot-latest.zip && \
	mv 0x_ganache_snapshot $hdir/ganache

echo "Done getting snapshot into $hdir"
