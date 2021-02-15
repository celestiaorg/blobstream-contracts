#!/bin/bash
yarn ts-node \
contract-deployer.ts \
--cosmos-node="http://localhost:26657" \
--eth-node="http://localhost:8545" \
--eth-privkey="0xf2f48ee19680706196e2e339e5da3491186e0c4c5030670656b0e0164837257d" \
--contract=artifacts/contracts/Peggy.sol/Peggy.json \
--test-mode=true
