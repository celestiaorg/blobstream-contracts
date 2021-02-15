#!/bin/bash

set -e

TM_NODE=$(grep PEGGY_TENDERMINT_RPC ../../orchestrator/.env | cut -d '=' -f2)
PASSPHRASE=$(grep PASSPHRASE ../../orchestrator/.env | cut -d '=' -f2)
PROPOSAL_ID=$(echo $(yes $PASSPHRASE | injectived tx gov submit-proposal param-change bridgeParamProposal.json --from genesis --chain-id=888 --keyring-backend=file --yes --node=$TM_NODE | jq '.logs[0].events[2].attributes[0].value') | sed 's/^"\(.*\)"$/\1/')

echo "Voting YES for Proposal" $PROPOSAL_ID

yes $PASSPHRASE | injectived tx gov vote $PROPOSAL_ID yes --from genesis --chain-id=888 --keyring-backend=file --yes --node=$TM_NODE