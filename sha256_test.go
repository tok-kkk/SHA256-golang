package sha256

import (
	"crypto/sha256"
	"testing/quick"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("sha256", func() {
	Context("sha256 hashing ", func() {
		It("should return the right hash", func() {
			// input := "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq"
			// log.Println(hex.EncodeToString(Sha256([]byte(input))))
			Expect(quick.CheckEqual(Sha256, sha256.Sum256, nil)).NotTo(HaveOccurred())

			Expect(true).Should(BeTrue())
		})
	})
})
