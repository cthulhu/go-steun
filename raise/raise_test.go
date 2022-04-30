package raise_test

import (
	"fmt"

	. "github.com/cthulhu/go-steun/raise"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func testF() {
	err := fmt.Errorf("error")
	PanicOnError(err)
}

var _ = Describe("Raise", func() {
	Context("Panic", func() {
		It("Panics if error", func() {
			Expect(testF).Should(Panic())
		})
	})
})
