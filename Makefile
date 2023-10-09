.PHONY: gen

###############################################################################
##                                 Solidity                                  ##
###############################################################################

gen: solidity-wrappers

SOLIDITY_DIR = .
SOLIDITY_SRC_DIR = $(SOLIDITY_DIR)/src
CONTRACTS = BlobStream.sol ERC1967Proxy.sol

fmt:
	@echo "--> Running forge fmt"
	@forge fmt

solidity-wrappers:
	./scripts/gen.sh $(SOLIDITY_SRC_DIR) $(CONTRACTS)
