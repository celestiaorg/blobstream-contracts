package solidity

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/InjectiveLabs/evm-deploy-contract/deployer"
	. "github.com/onsi/ginkgo"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Returns a hex string representation of text, exactly 32 bytes wide. Strings must be 31 bytes or shorter, or an exception is thrown.
//
// NOTE: Keep in mind that UTF-8 characters outside the ASCII range can be multiple bytes long.
func formatBytes32String(str string) common.Hash {
	var v common.Hash
	copy(v[:], str)
	return v
}

func makeCheckpoint(
	validators []common.Address,
	powers []*big.Int,
	valsetNonce *big.Int,
	peggyId common.Hash,
) common.Hash {
	methodName := formatBytes32String("checkpoint")

	buf, err := checkpointABI.Pack("checkpoint",
		peggyId, methodName, valsetNonce, validators, powers,
	)
	orFail(err)

	return crypto.Keccak256Hash(buf[4:])
}

var checkpointABI, _ = abi.JSON(strings.NewReader(checkpointABIJSON))

var checkpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "pure",
		"type": "function",
		"inputs": [
			{ "internalType": "bytes32",   "name": "_peggyId",     "type": "bytes32" },
			{ "internalType": "bytes32",   "name": "_methodName",  "type": "bytes32" },
			{ "internalType": "uint256",   "name": "_valsetNonce", "type": "uint256" },
			{ "internalType": "address[]", "name": "_validators",  "type": "address[]" },
			{ "internalType": "uint256[]", "name": "_powers",      "type": "uint256[]" }
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

var zeroAddress = common.Address{}

var zeroHash = common.Hash{}

// maxUInt256 returns a value equal to 2**256 - 1 (MAX_UINT in Solidity).
func maxUInt256() *big.Int {
	return new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
}

func orFail(err error) {
	if err != nil {
		Fail(err.Error(), 1)
	}
}

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}

var noArgs = func(args abi.Arguments) []interface{} {
	return nil
}

func withArgsFn(args ...interface{}) deployer.AbiMethodInputMapperFunc {
	return func(_ abi.Arguments) []interface{} {
		return args
	}
}

func sumInts(n0 *big.Int, n ...*big.Int) *big.Int {
	sum := new(big.Int)
	if n0 != nil {
		sum.Set(n0)
	}

	for _, i := range n {
		sum.Add(sum, i)
	}

	return sum
}

func signDigest(digestHash common.Hash, keys ...*ecdsa.PrivateKey) (
	v []uint8,
	r []common.Hash,
	s []common.Hash,
	err error,
) {
	// The produced signature is in the [R || S || V] format where V is 0 or 1.
	var sig []byte

	for _, k := range keys {
		sig, err = crypto.Sign(digestHash[:], k)
		if err != nil {
			return
		}

		sigV := sig[64]
		sigR := common.Hash{}
		_ = copy(sigR[:], sig[:32])
		sigS := common.Hash{}
		_ = copy(sigS[:], sig[32:64])

		v = append(v, sigV)
		r = append(r, sigR)
		s = append(s, sigS)
	}

	return
}
