package cosmos

import (
	"context"
	"math/big"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/umee-network/peggo/mocks"
	wrappers "github.com/umee-network/peggo/solidity/wrappers/Peggy.sol"
	"github.com/umee-network/umee/x/peggy/types"
)

type hasBiggerNonce struct {
	currentNonce uint64
}

func (m *hasBiggerNonce) Matches(input interface{}) bool {
	deposit, ok := input.(*types.MsgDepositClaim)
	if ok {
		if deposit.EventNonce > m.currentNonce {
			m.currentNonce = deposit.EventNonce

			return true
		}
		return false
	}

	withdraw, ok := input.(*types.MsgWithdrawClaim)
	if ok {
		if withdraw.EventNonce > m.currentNonce {
			m.currentNonce = withdraw.EventNonce
			return true
		}
	}

	valsetUpdate, ok := input.(*types.MsgValsetUpdatedClaim)
	if ok {
		if valsetUpdate.EventNonce > m.currentNonce {
			m.currentNonce = valsetUpdate.EventNonce
			return true
		}
	}

	erc20Deployed, ok := input.(*types.MsgERC20DeployedClaim)
	if ok {
		if erc20Deployed.EventNonce > m.currentNonce {
			m.currentNonce = erc20Deployed.EventNonce
			return true
		}
	}

	return false
}

func (m *hasBiggerNonce) String() string {
	return "nonce must be higher"
}

func HasBiggerNonce(initialNonce uint64) gomock.Matcher {
	return &hasBiggerNonce{currentNonce: initialNonce}
}

func TestSendEthereumClaims(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCosmos := mocks.NewMockCosmosClient(mockCtrl)
	mockCosmos.EXPECT().FromAddress().Return(sdk.AccAddress{}).AnyTimes()

	biggerNonceMatcher := HasBiggerNonce(0)
	mockCosmos.EXPECT().SyncBroadcastMsg(biggerNonceMatcher).Return(&sdk.TxResponse{}, nil).Times(8)

	s := peggyBroadcastClient{
		daemonQueryClient: nil,
		broadcastClient:   mockCosmos,
	}

	deposits := []*wrappers.PeggySendToCosmosEvent{
		{
			EventNonce: big.NewInt(2),
			Amount:     big.NewInt(123),
		},
		{
			EventNonce: big.NewInt(6),
			Amount:     big.NewInt(456),
		},
	}

	withdraws := []*wrappers.PeggyTransactionBatchExecutedEvent{
		{
			EventNonce: big.NewInt(1),
			BatchNonce: big.NewInt(0),
		},
		{
			EventNonce: big.NewInt(3),
			BatchNonce: big.NewInt(0),
		},
	}

	valsetUpdates := []*wrappers.PeggyValsetUpdatedEvent{
		{
			EventNonce:     big.NewInt(4),
			NewValsetNonce: big.NewInt(0),
			RewardAmount:   big.NewInt(0),
		},
		{
			EventNonce:     big.NewInt(5),
			NewValsetNonce: big.NewInt(0),
			RewardAmount:   big.NewInt(0),
		},
		{
			EventNonce:     big.NewInt(7),
			NewValsetNonce: big.NewInt(0),
			RewardAmount:   big.NewInt(0),
		},
	}

	erc20Deployed := []*wrappers.PeggyERC20DeployedEvent{
		{
			EventNonce: big.NewInt(8),
		},
	}

	s.SendEthereumClaims(context.Background(),
		0,
		deposits,
		withdraws,
		valsetUpdates,
		erc20Deployed,
		time.Microsecond,
	)
}

func TestSendEthereumClaimsIgnoreNonSequentialNonces(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCosmos := mocks.NewMockCosmosClient(mockCtrl)
	mockCosmos.EXPECT().FromAddress().Return(sdk.AccAddress{}).AnyTimes()

	biggerNonceMatcher := HasBiggerNonce(0)
	mockCosmos.EXPECT().SyncBroadcastMsg(biggerNonceMatcher).Return(&sdk.TxResponse{}, nil).Times(7)

	s := peggyBroadcastClient{
		daemonQueryClient: nil,
		broadcastClient:   mockCosmos,
	}

	// We have events with nonces 1, 2, 3, 4, 5, 6, 7, 9.
	// So we are missing the 8, meaning events above that won't be relayed
	deposits := []*wrappers.PeggySendToCosmosEvent{
		{
			EventNonce: big.NewInt(2),
			Amount:     big.NewInt(123),
		},
		{
			EventNonce: big.NewInt(6),
			Amount:     big.NewInt(456),
		},
	}

	withdraws := []*wrappers.PeggyTransactionBatchExecutedEvent{
		{
			EventNonce: big.NewInt(1),
			BatchNonce: big.NewInt(0),
		},
		{
			EventNonce: big.NewInt(3),
			BatchNonce: big.NewInt(0),
		},
	}

	valsetUpdates := []*wrappers.PeggyValsetUpdatedEvent{
		{
			EventNonce:     big.NewInt(4),
			NewValsetNonce: big.NewInt(0),
			RewardAmount:   big.NewInt(0),
		},
		{
			EventNonce:     big.NewInt(5),
			NewValsetNonce: big.NewInt(0),
			RewardAmount:   big.NewInt(0),
		},
		{
			EventNonce:     big.NewInt(9),
			NewValsetNonce: big.NewInt(0),
			RewardAmount:   big.NewInt(0),
		},
	}

	erc20Deployed := []*wrappers.PeggyERC20DeployedEvent{
		{
			EventNonce: big.NewInt(7),
		},
	}

	s.SendEthereumClaims(context.Background(),
		0,
		deposits,
		withdraws,
		valsetUpdates,
		erc20Deployed,
		time.Microsecond,
	)
}
