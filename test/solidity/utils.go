package solidity

import (
	. "github.com/onsi/ginkgo"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Returns a hex string representation of text, exactly 32 bytes wide. Strings must be 31 bytes or shorter, or an exception is thrown.
//
// NOTE: Keep in mind that UTF-8 characters outside the ASCII range can be multiple bytes long.
func formatBytes32String(str string) common.Hash {
	var v common.Hash
	copy(v[:], str)
	return v
}

var zeroAddress = common.Address{}

func orFail(err error) {
	if err != nil {
		Fail(err.Error(), 1)
	}
}

var noArgs = func(args abi.Arguments) []interface{} {
	return nil
}
