package solidity

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/InjectiveLabs/evm-deploy-contract/deployer"
)

var _ = Describe("Hashing Test", func() {
	var (
		hashingTestContract common.Address
	)

	BeforeEach(func() {
		d, err := deployer.New()
		orFail(err)

		deployOpts := deployer.ContractDeployOpts{
			EVMEndpoint:  os.Getenv("SOLIDITY_TEST_EVM_RPC"),
			From:         EthAccounts[0].EthAddress,
			FromPk:       EthAccounts[0].EthPrivKey,
			SolSource:    "../../solidity/contracts/HashingTest.sol",
			ContractName: "HashingTest",
		}

		_, contract, err := d.Deploy(context.Background(), deployOpts, noArgs, false, true)
		orFail(err)

		hashingTestContract = contract.Address
	})

	_ = Context("HashingTest contract deployed", func() {
		It("Should have address", func() {
			Ω(hashingTestContract).ShouldNot(Equal(zeroAddress))
		})
	})

	// var (
	//     longBook  Book
	//     shortBook Book
	// )

	// BeforeEach(func() {
	//     longBook = Book{
	//         Title:  "Les Miserables",
	//         Author: "Victor Hugo",
	//         Pages:  1488,
	//     }

	//     shortBook = Book{
	//         Title:  "Fox In Socks",
	//         Author: "Dr. Seuss",
	//         Pages:  24,
	//     }
	// })

	// Describe("Categorizing book length", func() {
	// 	Context("With more than 300 pages", func() {
	// 		It("should be a novel", func() {
	// 			Ω(nil)
	// 			// Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
	// 		})
	// 	})

	// 	Context("With fewer than 300 pages", func() {
	// 		It("should be a short story", func() {
	// 			Ω(nil)
	// 			// Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
	// 		})
	// 	})
	// })
})
