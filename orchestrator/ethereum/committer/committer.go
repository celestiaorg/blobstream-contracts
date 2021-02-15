package committer

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/provider"
)

// EVMCommitter defines an interface for submitting transactions
// into Ethereum, Matic, and other EVM-compatible networks.
type EVMCommitter interface {
	FromAddress() common.Address
	Provider() provider.EVMProvider
	SendTx(recipient common.Address, txData []byte) (txHash common.Hash, err error)
}
