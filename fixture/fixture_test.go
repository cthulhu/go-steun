package fixture_test

import (
	. "github.com/cthulhu/go-steun/fixture"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fixture", func() {
	Context("Read", func() {
		It("reads the fixture", func() {
			Expect(string(Read("one.txt"))).To(BeEquivalentTo("expected\n"))
		})
	})
	Context("ReadFile", func() {
		It("reads the fixture", func() {
			Expect(string(ReadFile("fixtures/one.txt"))).To(BeEquivalentTo("expected\n"))
		})
	})
})
