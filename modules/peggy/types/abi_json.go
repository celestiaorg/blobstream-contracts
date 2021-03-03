package types

const (
	// OutgoingBatchTxCheckpointABIJSON checks the ETH ABI for compatability of the OutgoingBatchTx message
	OutgoingBatchTxCheckpointABIJSON = `[{
		"name": "submitBatch",
		"stateMutability": "pure",
		"type": "function",
		"inputs": [
			{ "internalType": "bytes32",   "name": "_peggyId",       "type": "bytes32" },
			{ "internalType": "bytes32",   "name": "_methodName",    "type": "bytes32" },
			{ "internalType": "uint256[]", "name": "_amounts",       "type": "uint256[]" },
			{ "internalType": "address[]", "name": "_destinations",  "type": "address[]" },
			{ "internalType": "uint256[]", "name": "_fees",          "type": "uint256[]" },
			{ "internalType": "uint256",   "name": "_batchNonce",    "type": "uint256" },
			{ "internalType": "address",   "name": "_tokenContract", "type": "address" }
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

	// ValsetCheckpointABIJSON checks the ETH ABI for compatability of the Valset update message
	ValsetCheckpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "pure",
		"type": "function",
		"inputs": [
			{ "internalType": "bytes32",   "name": "_peggyId",     "type": "bytes32" },
			{ "internalType": "bytes32",   "name": "_checkpoint",  "type": "bytes32" },
			{ "internalType": "uint256",   "name": "_valsetNonce", "type": "uint256" },
			{ "internalType": "address[]", "name": "_validators",  "type": "address[]" },
			{ "internalType": "uint256[]", "name": "_powers",      "type": "uint256[]" }
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`
)
