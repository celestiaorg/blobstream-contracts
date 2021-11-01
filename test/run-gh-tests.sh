#!/bin/bash

npm install -g ganache-cli
ganache-cli -h "0.0.0.0" -p 8545 -m "concert load couple harbor equip island argue ramp clarify fence smart topic" -l 999999999999999 &

wget https://binaries.soliditylang.org/linux-amd64/solc-linux-amd64-v0.8.2+commit.661d1103
mv solc-linux-amd64-v0.8.2+commit.661d1103 solc && chmod +x solc && mv solc /usr/local/bin/solc
solc --version

git clone --depth 1 --branch v0.4.0-rc2 https://github.com/umee-network/umee.git
cd umee
make install
cd ..

CHAIN_ID=umee-local STAKE_DENOM=uumee DENOM=uumee CLEANUP=1 ./test/cosmos/multinode.sh umeed
PEGGO_TEST_EVM_RPC="http://0.0.0.0:8545" go test ./...
