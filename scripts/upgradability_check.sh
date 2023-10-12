#!/usr/bin/env bash

# this script will check if the Blobstream contract is inheriting the correct upgradability contracts.

out=$(surya inheritance src/Blobstream.sol | grep -i "\"Blobstream\" ->" | cut -d ">" -f 2  | sed 's/[";]//g')

required_contracts=("Initializable" "UUPSUpgradeable" "OwnableUpgradeable")
missing_contracts=()

for field in "${required_contracts[@]}"; do
    if ! grep -q "\<$field\>" <<< "$out"; then
        missing_contracts+=("$field")
    fi
done

if [ ${#missing_contracts[@]} -eq 0 ]; then
    echo "The Blobstream contract is inheriting the right contracts. Exiting."
    exit 0
else
    echo "The Blobstream contract is missing the necessary inherited contracts: ${missing_contracts[*]}"
    exit 1
fi
