package structs_test

import (
	. "github.com/Shop2market/bonobo/support/structs"
	"github.com/Shop2market/yunnan/support"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestStruct struct {
	Name   string
	Age    int64
	Weight float64
	Job    *string
}

var _ = Describe("Structs", func() {
	Context("Plain structs", func() {
		It("Sets string values", func() {
			testStruct := &TestStruct{}
			Set(testStruct, "Name", "John")
			Expect(testStruct.Name).To(BeEquivalentTo("John"))
		})
		It("Sets int64 values", func() {
			testStruct := &TestStruct{}
			var age int64 = 20
			Expect(Set(testStruct, "Age", age)).NotTo(HaveOccurred())
			Expect(testStruct.Age).To(BeEquivalentTo(20))
		})
		It("Sets float64 values", func() {
			testStruct := &TestStruct{}
			var weight float64 = 75.5
			Expect(Set(testStruct, "Weight", weight)).NotTo(HaveOccurred())
			Expect(testStruct.Weight).To(BeEquivalentTo(75.5))
		})
		It("Sets *string values", func() {
			testStruct := &TestStruct{}
			var job string = "IT"
			Expect(Set(testStruct, "Job", &job)).NotTo(HaveOccurred())
			Expect(testStruct.Job).To(BeEquivalentTo(support.StrPtr("IT")))
		})

	})
})
