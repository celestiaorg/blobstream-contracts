#!/usr/bin/env bash

set -e

if (( $# < 2 )); then
    echo "Go wrappers generator script. Make sure to specify the following params:"
    echo " - first parameter: the contracts source directory"
    echo " - second parameter: the contracts names (including the .sol extension) separated by a space"
    echo "the output files will be in the ./wrappers directory."
    exit 1
fi

forge build > /dev/null

cd "$1"

for file in "${@: 2}"; do
    mkdir -p ../wrappers/"${file}"
    contractName=$(basename "${file}" .sol)

    jq .abi < ../out/"${file}"/"${contractName}".json > ../out/"${file}"/"${contractName}".abi
    jq -r .bytecode.object < ../out/"${file}"/"${contractName}".json > ../out/"${file}"/"${contractName}".bin

    abigen --type=qgb --pkg wrappers \
        --out=../wrappers/"${file}"/wrapper.go \
        --abi ../out/"${file}"/"${contractName}".abi \
        --bin ../out/"${file}"/"${contractName}".bin
done

echo "done."
