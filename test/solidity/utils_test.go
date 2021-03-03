package solidity

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing utils check", func() {
	_ = Describe("Test formatBytes32String", func() {
		It("Should format string of length < 32 bytes", func() {
			result := formatBytes32String("Hello World!").Hex()
			Ω(result).Should(Equal("0x48656c6c6f20576f726c64210000000000000000000000000000000000000000"))

			result = formatBytes32String("checkpoint").Hex()
			Ω(result).Should(Equal("0x636865636b706f696e7400000000000000000000000000000000000000000000"))
		})
	})
})
