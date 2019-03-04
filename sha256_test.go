package sha256

import (
	"crypto/sha256"
	"testing"
	"testing/quick"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("sha256", func() {
	Context("sha256 hashing ", func() {
		It("should return the right hash", func() {
			Expect(quick.CheckEqual(Sha256, sha256.Sum256, nil)).NotTo(HaveOccurred())
		})
	})
})

func BenchmarkSha256(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sha256([]byte("abcdefghijklmnopqrstuvwxyz"))
	}
}

func BenchmarkStandSha256(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		sha256.Sum256([]byte("abcdefghijklmnopqrstuvwxyz"))
	}
}

