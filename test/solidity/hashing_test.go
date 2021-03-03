package solidity

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/InjectiveLabs/evm-deploy-contract/deployer"
)

var _ = Describe("Contract Tests", func() {
	_ = Context("Hashing Test", func() {
		var (
			hashingTestTxOpts   deployer.ContractTxOpts
			hashingTestCallOpts deployer.ContractCallOpts
			hashingTestContract common.Address
		)

		_ = It("Deploys HashingTest.sol", func() {
			hashingTestDeployOpts := deployer.ContractDeployOpts{
				From:         EthAccounts[0].EthAddress,
				FromPk:       EthAccounts[0].EthPrivKey,
				SolSource:    "../../solidity/contracts/HashingTest.sol",
				ContractName: "HashingTest",
				Await:        true,
			}

			_, contract, err := ContractDeployer.Deploy(context.Background(), hashingTestDeployOpts, noArgs)
			Ω(err).Should(BeNil())
			Ω(contract.Address).ShouldNot(Equal(zeroAddress))

			hashingTestContract = contract.Address
		})

		BeforeEach(func() {
			hashingTestTxOpts = deployer.ContractTxOpts{
				From:         EthAccounts[0].EthAddress,
				FromPk:       EthAccounts[0].EthPrivKey,
				SolSource:    "../../solidity/contracts/HashingTest.sol",
				ContractName: "HashingTest",
				Contract:     hashingTestContract,
				Await:        true,
			}

			hashingTestCallOpts = deployer.ContractCallOpts{
				From:         EthAccounts[0].EthAddress,
				SolSource:    "../../solidity/contracts/HashingTest.sol",
				ContractName: "HashingTest",
				Contract:     hashingTestContract,
			}
		})

		_ = When("HashingTest contract deployment done", func() {
			var (
				peggyID     common.Hash
				validators  []common.Address
				powers      []*big.Int
				valsetNonce *big.Int
			)

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

			It("Update checkpoint using IterativeHash", func() {
				_, _, err := ContractDeployer.Tx(context.Background(), hashingTestTxOpts,
					"IterativeHash", withArgsFn(validators, powers, valsetNonce, peggyID),
				)
				Ω(err).Should(BeNil())
			})

			It("Update checkpoint using ConcatHash", func() {
				_, _, err := ContractDeployer.Tx(context.Background(), hashingTestTxOpts,
					"ConcatHash", withArgsFn(validators, powers, valsetNonce, peggyID),
				)
				Ω(err).Should(BeNil())
			})

			It("Update checkpoint using ConcatHash2", func() {
				_, _, err := ContractDeployer.Tx(context.Background(), hashingTestTxOpts,
					"ConcatHash2", withArgsFn(validators, powers, valsetNonce, peggyID),
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
					makeCheckpoint(validators, powers, valsetNonce, peggyID),
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
