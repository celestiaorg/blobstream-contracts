package solidity

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ctypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/InjectiveLabs/evm-deploy-contract/deployer"

	"github.com/InjectiveLabs/evm-deploy-contract/sol"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/peggy"
	wrappers "github.com/InjectiveLabs/peggo/solidity/wrappers/Peggy.sol"
)

var _ = Describe("Contract Tests", func() {
	_ = Describe("Peggy", func() {
		var (
			peggyTxOpts   deployer.ContractTxOpts
			peggyCallOpts deployer.ContractCallOpts
			peggyLogsOpts deployer.ContractLogsOpts
			peggyContract *sol.Contract

			deployArgs   deployer.AbiMethodInputMapperFunc
			deployErr    error
			deployTxHash common.Hash

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

			deployTxHash, peggyContract, deployErr = ContractDeployer.Deploy(context.Background(), peggyDeployOpts, deployArgs)
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
				peggyOwner Account
			)

			BeforeEach(func() {
				peggyOwner = EthAccounts[0]
			})

			JustBeforeEach(func() {
				orFail(deployErr)

				peggyTxOpts = deployer.ContractTxOpts{
					From:         peggyOwner.EthAddress,
					FromPk:       peggyOwner.EthPrivKey,
					SolSource:    "../../solidity/contracts/Peggy.sol",
					ContractName: "Peggy",
					Contract:     peggyContract.Address,
					Await:        true,
				}

				peggyCallOpts = deployer.ContractCallOpts{
					From:         peggyOwner.EthAddress,
					SolSource:    "../../solidity/contracts/Peggy.sol",
					ContractName: "Peggy",
					Contract:     peggyContract.Address,
				}

				peggyLogsOpts = deployer.ContractLogsOpts{
					SolSource:    "../../solidity/contracts/Peggy.sol",
					ContractName: "Peggy",
					Contract:     peggyContract.Address,
				}
			})

			_ = Describe("Check contract state", func() {
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

				It("Should have generated a valid checkpoint", func() {
					var state_lastValsetCheckpoint common.Hash

					out, outAbi, err := ContractDeployer.Call(context.Background(), peggyCallOpts,
						"state_lastValsetCheckpoint", noArgs,
					)
					Ω(err).Should(BeNil())

					offchainCheckpoint := makeCheckpoint(validators, powers, big.NewInt(0), peggyID)

					err = outAbi.Copy(&state_lastValsetCheckpoint, out)
					Ω(err).Should(BeNil())
					Ω(state_lastValsetCheckpoint).Should(Equal(offchainCheckpoint))
				})

				_ = Describe("ValsetUpdatedEvent", func() {
					var (
						valsetUpdatedEvent = wrappers.PeggyValsetUpdatedEvent{}
					)

					BeforeEach(func() {
						_, err := ContractDeployer.Logs(
							context.Background(),
							peggyLogsOpts,
							deployTxHash,
							"ValsetUpdatedEvent",
							unpackValsetUpdatedEventTo(&valsetUpdatedEvent),
						)
						orFail(err)
					})

					It("Should have valid Valset parameters", func() {
						Ω(valsetUpdatedEvent.NewValsetNonce).ShouldNot(BeNil())
						Ω(valsetUpdatedEvent.NewValsetNonce.String()).Should(Equal("0"))
						Ω(valsetUpdatedEvent.Validators).Should(BeEquivalentTo(validators))
						Ω(valsetUpdatedEvent.Powers).Should(BeEquivalentTo(powers))
					})
				})
			})

			_ = Describe("ERC20 token deployment via Peggy", func() {
				var (
					state_lastEventNonce *big.Int
					prevEventNonce       *big.Int

					erc20DeployTxHash  common.Hash
					erc20DeployErr     error
					erc20DeployedEvent = wrappers.PeggyERC20DeployedEvent{}
				)

				BeforeEach(func() {
					if state_lastEventNonce != nil {
						prevEventNonce = state_lastEventNonce
					}

					out, outAbi, err := ContractDeployer.Call(context.Background(), peggyCallOpts,
						"state_lastEventNonce", noArgs,
					)
					Ω(err).Should(BeNil())
					err = outAbi.Copy(&state_lastEventNonce, out)
					Ω(err).Should(BeNil())
				})

				It("Deploys a new ERC20 contract instance", func() {
					erc20DeployTxHash, _, erc20DeployErr = ContractDeployer.Tx(context.Background(), peggyTxOpts,
						"deployERC20", withArgsFn("inj", "INJ", "INJ", byte(18)),
					)
					Ω(erc20DeployErr).Should(BeNil())
					Ω(erc20DeployTxHash).ShouldNot(Equal(zeroHash))
				})

				It("Nonce during deployment increased", func() {
					next := new(big.Int).Add(prevEventNonce, big.NewInt(1))
					Ω(state_lastEventNonce.String()).Should(Equal(next.String()))
				})

				_ = When("New ERC20 instance deployed", func() {
					BeforeEach(func() {
						orFail(erc20DeployErr)

						_, err := ContractDeployer.Logs(
							context.Background(),
							peggyLogsOpts,
							erc20DeployTxHash,
							"ERC20DeployedEvent",
							unpackERC20DeployedEventTo(&erc20DeployedEvent),
						)
						orFail(err)
					})

					_ = Describe("ERC20DeployedEvent", func() {
						It("Should have valid token params", func() {
							Ω(erc20DeployedEvent.CosmosDenom).Should(Equal("inj"))
							Ω(erc20DeployedEvent.Symbol).Should(Equal("INJ"))
							Ω(erc20DeployedEvent.Name).Should(Equal("INJ"))
							Ω(erc20DeployedEvent.Decimals).Should(BeEquivalentTo(18))
						})

						It("Should have TokenContract address", func() {
							Ω(erc20DeployedEvent.TokenContract).ShouldNot(Equal(zeroAddress))
						})

						It("Should have valid EventNonce", func() {
							Ω(erc20DeployedEvent.EventNonce).ShouldNot(BeNil())
							Ω(erc20DeployedEvent.EventNonce.String()).Should(Equal(state_lastEventNonce.String()))
						})
					})

					_ = Describe("ERC20 Token", func() {
						var (
							erc20CallOpts deployer.ContractCallOpts
						)

						BeforeEach(func() {
							erc20CallOpts = deployer.ContractCallOpts{
								From:         peggyOwner.EthAddress,
								SolSource:    "../../solidity/contracts/CosmosToken.sol",
								ContractName: "CosmosERC20",
								Contract:     erc20DeployedEvent.TokenContract,
							}
						})

						It("Should have MAX_UINT balance on Peggy", func() {
							var peggyBalance *big.Int

							out, outAbi, err := ContractDeployer.Call(context.Background(), erc20CallOpts,
								"balanceOf", withArgsFn(peggyContract.Address))
							Ω(err).Should(BeNil())

							err = outAbi.Copy(&peggyBalance, out)
							Ω(err).Should(BeNil())

							Ω(peggyBalance).Should(BeEquivalentTo(maxUInt256()))
						})

						_ = When("Cosmos -> Ethereum batch being submitted", func() {
							var (
								submitBatchTxHash common.Hash
								submitBatchErr    error
								prepareBatchErr   error
								signBatchErr      error

								txAmounts            []*big.Int
								txDestinations       []common.Address
								txFees               []*big.Int
								transactionBatchHash common.Hash

								sigsV []uint8
								sigsR []common.Hash
								sigsS []common.Hash

								currentValsetNonce *big.Int
								batchNonce         *big.Int
								batchTimeout       *big.Int
							)

							BeforeEach(func() {
								currentValsetNonce = big.NewInt(0)
								batchNonce = big.NewInt(1)
								batchTimeout = big.NewInt(10000)

								txAmounts = make([]*big.Int, len(EthAccounts))
								txDestinations = getEthAddresses(EthAccounts...)
								txFees = make([]*big.Int, len(EthAccounts))

								for i := range EthAccounts {
									txAmounts[i] = big.NewInt(1)
									txFees[i] = big.NewInt(1)
								}

								transactionBatchHash, prepareBatchErr = prepareOutgoingTransferBatch(
									peggyID,
									erc20DeployedEvent.TokenContract,
									txAmounts,
									txDestinations,
									txFees,
									batchNonce,
								)
								orFail(prepareBatchErr)

								sigsV, sigsR, sigsS, signBatchErr = signDigest(
									transactionBatchHash, getSigningKeys(CosmosAccounts[:3]...)...)
								orFail(signBatchErr)
							})

							JustBeforeEach(func() {
								// don't resend the batch
								if submitBatchTxHash != zeroHash {
									return
								}

								submitBatchTxHash, _, submitBatchErr = ContractDeployer.Tx(context.Background(), peggyTxOpts,
									"submitBatch", withArgsFn(
										// The validators that approve the batch
										validators,         // 	address[] memory _currentValidators,
										powers,             // 	uint256[] memory _currentPowers,
										currentValsetNonce, // 	uint256 _currentValsetNonce,

										// These are arrays of the parts of the validators signatures
										sigsV, // 	uint8[] memory _v,
										sigsR, // 	bytes32[] memory _r,
										sigsS, // 	bytes32[] memory _s,

										// The batch of transactions
										txAmounts,                        // 	uint256[] memory _amounts,
										txDestinations,                   // 	address[] memory _destinations,
										txFees,                           // 	uint256[] memory _fees,
										batchNonce,                       // 	uint256 _batchNonce,
										erc20DeployedEvent.TokenContract, // 	address _tokenContract,

										// a block height beyond which this batch is not valid
										// used to provide a fee-free timeout
										batchTimeout, // 	uint256 _batchTimeout
									))

							})

							_ = When("TxBatch submission failed", func() {
								BeforeEach(func() {})
								AfterEach(func() {})
							})

							_ = Context("TxBatch submitted successfully", func() {
								BeforeEach(func() {
									orFail(submitBatchErr)
								})

								It("Changes the balance of the Peggy contract", func() {
									var peggyBalance *big.Int

									out, outAbi, err := ContractDeployer.Call(context.Background(), erc20CallOpts,
										"balanceOf", withArgsFn(peggyContract.Address))
									Ω(err).Should(BeNil())

									err = outAbi.Copy(&peggyBalance, out)
									Ω(err).Should(BeNil())

									expenses := sumInts(nil, txAmounts...)
									expenses = sumInts(expenses, txFees...)
									remainder := new(big.Int).Sub(maxUInt256(), expenses)
									Ω(peggyBalance.String()).Should(Equal(remainder.String()))
								})

								It("Increases the token balances of recipients", func() {
									for _, recipient := range getEthAddresses(EthAccounts...) {
										var recipientBalance *big.Int

										out, outAbi, err := ContractDeployer.Call(context.Background(), erc20CallOpts,
											"balanceOf", withArgsFn(recipient))
										Ω(err).Should(BeNil())

										err = outAbi.Copy(&recipientBalance, out)
										Ω(err).Should(BeNil())

										Ω(recipientBalance.String()).Should(Equal("1"))
									}
								})
							})
						})
					})
				})
			})
		})
	})
})

var outgoingBatchTxConfirmABI, _ = abi.JSON(strings.NewReader(peggy.OutgoingBatchTxConfirmABIJSON))

func prepareOutgoingTransferBatch(
	peggyID common.Hash,
	tokenContract common.Address,
	txAmounts []*big.Int,
	txDestinations []common.Address,
	txFees []*big.Int,
	batchNonce *big.Int,
) (common.Hash, error) {
	abiEncodedBatch, err := outgoingBatchTxConfirmABI.Pack("transactionBatch",
		peggyID,
		formatBytes32String("transactionBatch"),
		txAmounts,
		txDestinations,
		txFees,
		batchNonce,
		tokenContract,
	)
	if err != nil {
		return common.Hash{}, err
	}

	hash := crypto.Keccak256Hash(abiEncodedBatch[4:])
	return common.BytesToHash(hash.Bytes()), nil
}

func unpackERC20DeployedEventTo(result *wrappers.PeggyERC20DeployedEvent) deployer.ContractLogUnpackFunc {
	return func(unpacker deployer.LogUnpacker, event abi.Event, log ctypes.Log) (interface{}, error) {
		err := unpacker.UnpackLog(result, event.Name, log)
		return &result, err
	}
}

func unpackValsetUpdatedEventTo(result *wrappers.PeggyValsetUpdatedEvent) deployer.ContractLogUnpackFunc {
	return func(unpacker deployer.LogUnpacker, event abi.Event, log ctypes.Log) (interface{}, error) {
		err := unpacker.UnpackLog(result, event.Name, log)
		return &result, err
	}
}
