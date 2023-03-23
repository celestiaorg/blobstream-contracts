.PHONY: gen

###############################################################################
##                                 Solidity                                  ##
###############################################################################

gen: solidity-wrappers

SOLIDITY_DIR = .
SOLIDITY_SRC_DIR = $(SOLIDITY_DIR)/src
solidity-wrappers: $(SOLIDITY_SRC_DIR)/QuantumGravityBridge.sol $(SOLIDITY_SRC_DIR)/DAVerifier.sol
	cd $(SOLIDITY_SRC_DIR) ; \
	for file in $(^F) ; do \
			mkdir -p ../wrappers/$${file} ; \
			echo abigen --type=peggy --pkg wrappers --out=../wrappers/$${file}/wrapper.go --sol $${file} ; \
			abigen --type=peggy --pkg wrappers --out=../wrappers/$${file}/wrapper.go --sol $${file} ; \
	done
