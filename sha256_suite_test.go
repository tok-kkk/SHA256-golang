package sha256_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSha256(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sha256 Suite")
}
