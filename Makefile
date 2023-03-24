.PHONY: gen

###############################################################################
##                                 Solidity                                  ##
###############################################################################

gen: solidity-wrappers

SOLIDITY_DIR = .
SOLIDITY_SRC_DIR = $(SOLIDITY_DIR)/src
CONTRACTS = QuantumGravityBridge.sol
solidity-wrappers:
	SOLIDITY_SRC_DIR=$(SOLIDITY_SRC_DIR) CONTRACTS=$(CONTRACTS) /bin/bash ./scripts/gen.sh
