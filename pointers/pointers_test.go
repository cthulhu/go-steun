package pointers_test

import (
	. "github.com/cthulhu/go-steun/pointers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pointers", func() {
	Context("StrPtr", func() {
		It("returns pointer to string", func() {
			actual := "Actual"
			Expect(StrPtr(actual)).To(BeEquivalentTo(&actual))
		})
	})
	Context("BoolPtr", func() {
		It("returns pointer to bool", func() {
			actual := true
			Expect(BoolPtr(actual)).To(BeEquivalentTo(&actual))
		})
	})
	Context("Float64Ptr", func() {
		It("returns pointer to float", func() {
			actual := 12.12
			Expect(Float64Ptr(actual)).To(BeEquivalentTo(&actual))
		})
	})
	Context("IntPtr", func() {
		It("returns pointer to integer", func() {
			actual := 12
			Expect(IntPtr(actual)).To(BeEquivalentTo(&actual))
		})
	})
})
