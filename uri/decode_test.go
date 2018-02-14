package uri_test

import (
	"testing"

	. "github.com/cthulhu/go-steun/uri"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decode", func() {
	Context("URIDecodeComponent", func() {
		It("decodes", func() {
			Expect(URIDecodeComponent("123%2D321")).To(BeEquivalentTo("123-321"))
			Expect(URIDecodeComponent("123%2E321")).To(BeEquivalentTo("123.321"))
			Expect(URIDecodeComponent("123%23321")).To(BeEquivalentTo("123#321"))
			Expect(URIDecodeComponent("123%2F321")).To(BeEquivalentTo("123/321"))
		})
	})
})

// go test -bench=. -benchmem
func BenchmarkSimple(b *testing.B) {
	in := "123%2D321"
	for n := 0; n < b.N; n++ {
		_, _ = URIDecodeComponent(in)
	}
}

func BenchmarkMultiple(b *testing.B) {
	in := "123%2D321|123%2E321|123%23321|123%2F321"
	for n := 0; n < b.N; n++ {
		_, _ = URIDecodeComponent(in)
	}
}
