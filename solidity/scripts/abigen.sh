#!/bin/bash

set -e

NEW_UUID=$(cat /dev/urandom | LC_CTYPE=C tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

solc --abi ./contracts/Peggy.sol -o build_${NEW_UUID}
abigen --abi build_${NEW_UUID}/Peggy.abi --type Peggy --pkg wrappers > peggy.go
rm -rf build_${NEW_UUID}
rm -f Peggy.abi
gofmt -w peggy.go && echo "Peggy wrapper generated."
