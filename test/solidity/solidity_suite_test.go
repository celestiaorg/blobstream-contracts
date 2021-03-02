package solidity

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSolidity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Solidity Suite")
}
