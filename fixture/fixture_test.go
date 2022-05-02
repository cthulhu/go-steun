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
		It("panics if no file", func() {
			Expect(func() {
				ReadFile("fixtures/one_that_doesn_t_exists.txt")
			}).To(Panic())
		})
		It("panics if file is dir", func() {
			Expect(func() {
				ReadFile("fixtures")
			}).To(Panic())
		})
	})
	Context("Path", func() {
		It("returns the fixture path", func() {
			Expect(Path("one")).To(ContainSubstring("/fixtures/one"))
		})
	})
})
