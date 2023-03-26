.PHONY: gen

###############################################################################
##                                 Solidity                                  ##
###############################################################################

gen: solidity-wrappers

SOLIDITY_DIR = .
SOLIDITY_SRC_DIR = $(SOLIDITY_DIR)/src
CONTRACTS = QuantumGravityBridge.sol IDAOracle.sol
solidity-wrappers:
	./scripts/gen.sh $(SOLIDITY_SRC_DIR) $(CONTRACTS)
