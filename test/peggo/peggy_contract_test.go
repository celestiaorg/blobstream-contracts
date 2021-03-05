package solidity

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/InjectiveLabs/evm-deploy-contract/deployer"
	"github.com/InjectiveLabs/evm-deploy-contract/sol"
)

var _ = Describe("Contract Tests", func() {
	_ = Describe("Peggy", func() {
		var (
			peggyTxOpts   deployer.ContractTxOpts
			peggyCallOpts deployer.ContractCallOpts
			peggyContract *sol.Contract

			deployArgs deployer.AbiMethodInputMapperFunc
			deployErr  error

			peggyID    common.Hash
			minPower   *big.Int
			validators []common.Address
			powers     []*big.Int
		)

		BeforeEach(func() {
			peggyID = formatBytes32String("foo")
			validators = getEthAddresses(CosmosAccounts[:3]...)
			minPower = big.NewInt(3500)
			powers = []*big.Int{
				big.NewInt(3000),
				big.NewInt(1500),
				big.NewInt(500),
			}

			deployArgs = withArgsFn(
				peggyID,
				minPower,
				validators,
				powers,
			)
		})

		JustBeforeEach(func() {
			// don't redeploy if already deployed
			if peggyContract != nil {
				return
			}

			peggyDeployOpts := deployer.ContractDeployOpts{
				From:         EthAccounts[0].EthAddress,
				FromPk:       EthAccounts[0].EthPrivKey,
				SolSource:    "../../solidity/contracts/Peggy.sol",
				ContractName: "Peggy",
				Await:        true,
			}

			_, peggyContract, deployErr = ContractDeployer.Deploy(context.Background(), peggyDeployOpts, deployArgs)
		})

		_ = Context("Contract fails to deploy", func() {
			AfterEach(func() {
				deployArgs = withArgsFn(
					peggyID,
					minPower,
					validators,
					powers,
				)

				// force redeployment
				peggyContract = nil
			})

			_ = When("Throws on malformed valset", func() {
				BeforeEach(func() {
					deployArgs = withArgsFn(
						peggyID,
						minPower,
						validators,
						powers[:1], // only one
					)
				})

				It("Should fail with error", func() {
					Ω(deployErr).ShouldNot(BeNil())
					Ω(deployErr.Error()).Should(ContainSubstring("Malformed current validator set"))
				})
			})

			_ = When("Throws on insufficient power", func() {
				BeforeEach(func() {
					deployArgs = withArgsFn(
						peggyID,
						big.NewInt(10000),
						validators,
						powers,
					)
				})

				It("Should fail with error", func() {
					Ω(deployErr).ShouldNot(BeNil())
					Ω(deployErr.Error()).Should(ContainSubstring("Submitted validator set signatures do not have enough power"))
				})
			})
		})

		_ = Context("Peggy contract deployment done", func() {
			var (
				ethFrom Account
			)

			BeforeEach(func() {
				ethFrom = EthAccounts[0]
			})

			JustBeforeEach(func() {
				orFail(deployErr)

				peggyTxOpts = deployer.ContractTxOpts{
					From:         ethFrom.EthAddress,
					FromPk:       ethFrom.EthPrivKey,
					SolSource:    "../../solidity/contracts/Peggy.sol",
					ContractName: "Peggy",
					Contract:     peggyContract.Address,
					Await:        true,
				}

				peggyCallOpts = deployer.ContractCallOpts{
					From:         ethFrom.EthAddress,
					SolSource:    "../../solidity/contracts/Peggy.sol",
					ContractName: "Peggy",
					Contract:     peggyContract.Address,
				}
			})

			It("Should have address", func() {
				Ω(peggyTxOpts.Contract).ShouldNot(Equal(zeroAddress))
				Ω(peggyCallOpts.Contract).ShouldNot(Equal(zeroAddress))
			})

			It("Should have valid power threshold", func() {
				var state_powerThreshold *big.Int

				out, outAbi, err := ContractDeployer.Call(context.Background(), peggyCallOpts,
					"state_powerThreshold", noArgs,
				)
				Ω(err).Should(BeNil())

				err = outAbi.Copy(&state_powerThreshold, out)
				Ω(err).Should(BeNil())
				Ω(state_powerThreshold.String()).Should(Equal(minPower.String()))
			})

			It("Should have valid peggyId", func() {
				var state_peggyId common.Hash

				out, outAbi, err := ContractDeployer.Call(context.Background(), peggyCallOpts,
					"state_peggyId", noArgs,
				)
				Ω(err).Should(BeNil())

				err = outAbi.Copy(&state_peggyId, out)
				Ω(err).Should(BeNil())
				Ω(state_peggyId).Should(Equal(peggyID))
			})
		})
	})
})
