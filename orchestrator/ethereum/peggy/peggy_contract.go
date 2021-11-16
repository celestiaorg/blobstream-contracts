package peggy

import (
	"context"
	"math"
	"math/big"
	"strings"

	wrappers "github.com/celestiaorg/quantum-gravity-bridge/ethereum/solidity/wrappers/QuantumGravityBridge.sol"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/committer"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/provider"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
	"github.com/umee-network/umee/x/peggy/types"
)

type Contract interface {
	committer.EVMCommitter

	Address() common.Address

	SendTransactionBatch(
		ctx context.Context,
		currentValset *types.Valset,
		batch *types.OutgoingTxBatch,
		confirms []*types.MsgConfirmBatch,
	) (*common.Hash, error)

	SendEthValsetUpdate(
		ctx context.Context,
		oldValset *types.Valset,
		newValset *types.Valset,
		confirms []*types.MsgValsetConfirm,
	) (*common.Hash, error)

	GetValsetNonce(
		ctx context.Context,
		callerAddress common.Address,
	) (*big.Int, error)

	GetBridgeID(
		ctx context.Context,
		callerAddress common.Address,
	) (common.Hash, error)
}

func NewPeggyContract(
	logger zerolog.Logger,
	ethCommitter committer.EVMCommitter,
	peggyAddress common.Address,
) (Contract, error) {
	ethPeggy, err := wrappers.NewQuantumGravityBridge(peggyAddress, ethCommitter.Provider())
	if err != nil {
		return nil, err
	}

	svc := &peggyContract{
		logger:       logger.With().Str("module", "peggy_contract").Logger(),
		EVMCommitter: ethCommitter,
		peggyAddress: peggyAddress,
		ethPeggy:     ethPeggy,
	}

	return svc, nil
}

type peggyContract struct {
	logger zerolog.Logger
	committer.EVMCommitter

	ethProvider  provider.EVMProvider
	peggyAddress common.Address
	ethPeggy     *wrappers.QuantumGravityBridge
}

func (s *peggyContract) Address() common.Address {
	return s.peggyAddress
}

// maxUintAllowance is uint constant MAX_UINT = 2**256 - 1
var maxUintAllowance = big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))

var (
	peggyABI, _ = abi.JSON(strings.NewReader(wrappers.QuantumGravityBridgeMetaData.ABI))
)

func sigToVRS(sigHex string) (v uint8, r, s common.Hash) {
	signatureBytes := common.FromHex(sigHex)
	vParam := signatureBytes[64]
	if vParam == byte(0) {
		vParam = byte(27)
	} else if vParam == byte(1) {
		vParam = byte(28)
	}

	v = vParam
	r = common.BytesToHash(signatureBytes[0:32])
	s = common.BytesToHash(signatureBytes[32:64])

	return
}

// The total power in the Peggy bridge is normalized to u32 max every
// time a validator set is created. This value of up to u32 max is then
// stored in a i64 to prevent overflow during computation.
const totalPeggyPower int64 = math.MaxUint32

// peggyPowerToPercent takes in an amount of power in the peggy bridge, returns a percentage of total
func peggyPowerToPercent(total *big.Int) float32 {
	d := decimal.NewFromBigInt(total, 0)
	f, _ := d.Div(decimal.NewFromInt(totalPeggyPower)).Shift(2).Float64()
	return float32(f)
}

var ErrInsufficientVotingPowerToPass = errors.New("insufficient voting power")
