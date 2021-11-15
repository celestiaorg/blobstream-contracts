package peggy

import (
	"fmt"
	"math/big"
	"strings"

	log "github.com/xlab/suplog"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/umee-network/umee/x/peggy/types"
)

// EncodeValsetConfirm takes the required input data and produces the required signature to confirm a validator
// set update on the Peggy Ethereum contract. This value will then be signed before being
// submitted to Cosmos, verified, and then relayed to Ethereum
func EncodeValsetConfirm(peggyID common.Hash, valset *types.Valset) common.Hash {
	// error case here should not occur outside of testing since the above is a constant
	contractAbi, abiErr := abi.JSON(strings.NewReader(ValsetCheckpointABIJSON))
	if abiErr != nil {
		log.Fatalln("bad ABI constant")
	}

	checkpointBytes := []uint8("checkpoint")
	var checkpoint [32]uint8
	copy(checkpoint[:], checkpointBytes)

	memberAddresses := make([]common.Address, len(valset.Members))
	convertedPowers := make([]*big.Int, len(valset.Members))
	for i, m := range valset.Members {
		memberAddresses[i] = common.HexToAddress(m.EthereumAddress)
		convertedPowers[i] = big.NewInt(int64(m.Power))
	}

	rewardToken := common.HexToAddress(valset.RewardToken)

	if valset.RewardAmount.BigInt() == nil {
		// this must be programmer error
		panic("Invalid reward amount passed in valset GetCheckpoint!")
	}
	rewardAmount := valset.RewardAmount.BigInt()

	// the word 'checkpoint' needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	bytes, packErr := contractAbi.Pack(
		"checkpoint",
		peggyID,
		checkpoint,
		big.NewInt(int64(valset.Nonce)),
		memberAddresses,
		convertedPowers,
		rewardAmount,
		rewardToken,
	)
	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if packErr != nil {
		panic(fmt.Sprintf("Error packing checkpoint! %s/n", packErr))
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you where to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	hash := crypto.Keccak256Hash(bytes[4:])
	return hash
}

// EncodeTxBatchConfirm takes the required input data and produces the required signature to confirm a transaction
// batch on the Peggy Ethereum contract. This value will then be signed before being
// submitted to Cosmos, verified, and then relayed to Ethereum
func EncodeTxBatchConfirm(peggyID common.Hash, batch *types.OutgoingTxBatch) common.Hash {
	abi, err := abi.JSON(strings.NewReader(OutgoingBatchTxConfirmABIJSON))
	if err != nil {
		log.Fatalln("bad ABI constant")
	}

	// Create the methodName argument which salts the signature
	methodNameBytes := []uint8("transactionBatch")
	var batchMethodName [32]uint8
	copy(batchMethodName[:], methodNameBytes)

	// Run through the elements of the batch and serialize them
	txAmounts := make([]*big.Int, len(batch.Transactions))
	txDestinations := make([]common.Address, len(batch.Transactions))
	txFees := make([]*big.Int, len(batch.Transactions))
	for i, tx := range batch.Transactions {
		txAmounts[i] = tx.Erc20Token.Amount.BigInt()
		txDestinations[i] = common.HexToAddress(tx.DestAddress)
		txFees[i] = tx.Erc20Fee.Amount.BigInt()
	}

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abi.Pack("transactionBatch",
		peggyID,
		batchMethodName,
		txAmounts,
		txDestinations,
		txFees,
		big.NewInt(int64(batch.BatchNonce)),
		common.HexToAddress(batch.TokenContract),
		big.NewInt(int64(batch.BatchTimeout)),
	)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		log.WithError(err).Errorln("Error packing transactionBatch!")
		return common.Hash{}

	}

	hash := crypto.Keccak256Hash(abiEncodedBatch[4:])
	return common.BytesToHash(hash.Bytes())
}

const (
	// ValsetConfirmABIJSON = `[{
	//     "name": "checkpoint",
	//     "stateMutability": "pure",
	//     "type": "function",
	//     "inputs": [
	//         { "internalType": "bytes32",   "name": "_peggyId",     "type": "bytes32" },
	//         { "internalType": "bytes32",   "name": "_checkpoint",  "type": "bytes32" },
	//         { "internalType": "uint256",   "name": "_valsetNonce", "type": "uint256" },
	//         { "internalType": "address[]", "name": "_validators",  "type": "address[]" },
	//         { "internalType": "uint256[]", "name": "_powers",      "type": "uint256[]" }
	//     ]
	// }]`

	// ValsetCheckpointABIJSON checks the ETH ABI for compatibility of the Valset update message
	ValsetCheckpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "pure",
		"type": "function",
		"inputs": [
			{ "internalType": "bytes32",   "name": "_peggyId",   "type": "bytes32"   },
			{ "internalType": "bytes32",   "name": "_checkpoint",  "type": "bytes32"   },
			{ "internalType": "uint256",   "name": "_valsetNonce", "type": "uint256"   },
			{ "internalType": "address[]", "name": "_validators",  "type": "address[]" },
			{ "internalType": "uint256[]", "name": "_powers",      "type": "uint256[]" },
			{ "internalType": "uint256",   "name": "_rewardAmount","type": "uint256"   },
			{ "internalType": "address",   "name": "_rewardToken", "type": "address"   }
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

	OutgoingBatchTxConfirmABIJSON = `[{
        "name": "transactionBatch",
        "stateMutability": "pure",
        "type": "function",
        "inputs": [
            { "internalType": "bytes32",   "name": "_peggyId",       "type": "bytes32" },
            { "internalType": "bytes32",   "name": "_methodName",    "type": "bytes32" },
            { "internalType": "uint256[]", "name": "_amounts",       "type": "uint256[]" },
            { "internalType": "address[]", "name": "_destinations",  "type": "address[]" },
            { "internalType": "uint256[]", "name": "_fees",          "type": "uint256[]" },
            { "internalType": "uint256",   "name": "_batchNonce",    "type": "uint256" },
            { "internalType": "address",   "name": "_tokenContract", "type": "address" },
            { "internalType": "uint256",   "name": "_batchTimeout",  "type": "uint256" }
        ]
    }]`
)
