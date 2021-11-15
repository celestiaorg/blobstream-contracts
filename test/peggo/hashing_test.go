package solidity

import (
	"context"
	"math/big"
	"strings"

	"github.com/InjectiveLabs/etherman/deployer"
	"github.com/InjectiveLabs/etherman/sol"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/umee-network/peggo/orchestrator/ethereum/peggy"
)

var _ = Describe("Contract Tests", func() {
	_ = Describe("Hashing Test", func() {
		var (
			hashingTestTxOpts   deployer.ContractTxOpts
			hashingTestCallOpts deployer.ContractCallOpts
			hashingTestContract *sol.Contract
			deployErr           error
		)

		JustBeforeEach(func() {
			if hashingTestContract != nil {
				return
			}

			hashingTestDeployOpts := deployer.ContractDeployOpts{
				From:          EthAccounts[0].EthAddress,
				FromPk:        EthAccounts[0].EthPrivKey,
				SolSource:     "../../solidity/contracts/HashingTest.sol",
				ContractName:  "HashingTest",
				Await:         true,
				CoverageAgent: CoverageAgent,
			}

			_, hashingTestContract, deployErr = ContractDeployer.Deploy(context.Background(), hashingTestDeployOpts, noArgs)
		})

		_ = It("Deploys HashingTest.sol", func() {
			Ω(deployErr).Should(BeNil())
			Ω(hashingTestContract).ShouldNot(BeNil())
			Ω(hashingTestContract.Address).ShouldNot(Equal(zeroAddress))
		})

		_ = Context("HashingTest contract deployment done", func() {
			var (
				peggyID     common.Hash
				validators  []common.Address
				powers      []*big.Int
				valsetNonce *big.Int
			)

			BeforeEach(func() {
				orFail(deployErr)

				hashingTestTxOpts = deployer.ContractTxOpts{
					From:          EthAccounts[0].EthAddress,
					FromPk:        EthAccounts[0].EthPrivKey,
					SolSource:     "../../solidity/contracts/HashingTest.sol",
					ContractName:  "HashingTest",
					Contract:      hashingTestContract.Address,
					Await:         true,
					CoverageAgent: CoverageAgent,
				}

				hashingTestCallOpts = deployer.ContractCallOpts{
					From:          EthAccounts[0].EthAddress,
					SolSource:     "../../solidity/contracts/HashingTest.sol",
					ContractName:  "HashingTest",
					Contract:      hashingTestContract.Address,
					CoverageAgent: CoverageAgent,
					CoverageCall: deployer.ContractCoverageCallOpts{
						FromPk: EthAccounts[0].EthPrivKey,
					},
				}
			})

			BeforeEach(func() {
				peggyID = formatBytes32String("foo")
				validators = getEthAddresses(CosmosAccounts[:3]...)
				powers = make([]*big.Int, len(validators))
				for i := range powers {
					powers[i] = big.NewInt(5000)
				}

				valsetNonce = big.NewInt(1)
			})

			It("Should have address", func() {
				Ω(hashingTestTxOpts.Contract).ShouldNot(Equal(zeroAddress))
				Ω(hashingTestCallOpts.Contract).ShouldNot(Equal(zeroAddress))
			})

			It("Update checkpoint using CheckpointHash", func() {
				_, _, err := ContractDeployer.Tx(context.Background(), hashingTestTxOpts,
					"CheckpointHash", withArgsFn(validators, powers, valsetNonce, big.NewInt(0), zeroAddress, peggyID),
				)
				Ω(err).Should(BeNil())
			})

			It("Ensure that checkpoint equals the off-chain version", func() {
				var lastCheckpoint common.Hash

				out, outAbi, err := ContractDeployer.Call(context.Background(), hashingTestCallOpts,
					"lastCheckpoint", noArgs,
				)
				Ω(err).Should(BeNil())

				err = outAbi.Copy(&lastCheckpoint, out)
				Ω(err).Should(BeNil())

				Ω(lastCheckpoint).ShouldNot(Equal(zeroHash))
				Ω(lastCheckpoint).Should(Equal(
					makeValsetCheckpoint(peggyID, validators, powers, valsetNonce, big.NewInt(0), zeroAddress),
				))
			})

			It("Saves everything", func() {
				_, _, err := ContractDeployer.Tx(context.Background(), hashingTestTxOpts,
					"JustSaveEverything", withArgsFn(validators, powers, valsetNonce),
				)
				Ω(err).Should(BeNil())
			})

			It("Saves everything again", func() {
				_, _, err := ContractDeployer.Tx(context.Background(), hashingTestTxOpts,
					"JustSaveEverythingAgain", withArgsFn(validators, powers, valsetNonce),
				)
				Ω(err).Should(BeNil())
			})
		})
	})
})

var valsetConfirmABI, _ = abi.JSON(strings.NewReader(peggy.ValsetCheckpointABIJSON))

func makeValsetCheckpoint(
	peggyID common.Hash,
	validators []common.Address,
	powers []*big.Int,
	valsetNonce *big.Int,
	rewardAmount *big.Int,
	rewardToken common.Address,
) common.Hash {
	methodName := formatBytes32String("checkpoint")

	//TODO: check if we want to add a reward amount and a reward token here

	buf, err := valsetConfirmABI.Pack("checkpoint",
		peggyID,
		methodName,
		valsetNonce,
		validators,
		powers,
		rewardAmount,
		rewardToken,
	)

	orFail(err)

	return crypto.Keccak256Hash(buf[4:])
}
