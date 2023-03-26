#!/usr/bin/env bash

set -e

forge build > /dev/null

cd "${SOLIDITY_SRC_DIR}"

for file in "${CONTRACTS[@]}"; do
    mkdir -p ../wrappers/"${file}"
    contractName=$(echo "${file}" | cut -d . -f 1)

    jq .abi < ../out/"${file}"/"${contractName}".json > ../out/"${file}"/"${contractName}".abi
    jq .bytecode.object < ../out/"${file}"/"${contractName}".json | cut -d \" -f 2 > ../out/"${file}"/"${contractName}".bin

    abigen --type=qgb --pkg wrappers \
        --out=../wrappers/"${file}"/wrapper.go \
        --abi ../out/"${file}"/"${contractName}".abi \
        --bin ../out/"${file}"/"${contractName}".bin
done

echo "done."
