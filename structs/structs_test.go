package structs_test

import (
	"github.com/cthulhu/go-steun/pointers"
	. "github.com/cthulhu/go-steun/structs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestStruct struct {
	Name   string
	Age    int64
	Weight float64
	Job    *string
}

var _ = Describe("Structs", func() {
	Context("Set", func() {
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
			Expect(testStruct.Job).To(BeEquivalentTo(pointers.StrPtr("IT")))
		})
	})
	Context("Get", func() {
		It("Gets string values", func() {
			testStruct := &TestStruct{Name: "John"}
			actual, err := Get(testStruct, "Name")
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(BeEquivalentTo("John"))
		})
		It("Gets int64 values", func() {
			testStruct := &TestStruct{Age: 20}
			actual, err := Get(testStruct, "Age")
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(BeEquivalentTo(20))
		})
		It("Gets float64 values", func() {
			testStruct := &TestStruct{Weight: 75.7}
			actual, err := Get(testStruct, "Weight")
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(BeEquivalentTo(75.7))

		})
		It("Gets *string values", func() {
			job := "IT"
			testStruct := &TestStruct{Job: &job}
			actual, err := Get(testStruct, "Job")
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(BeEquivalentTo(pointers.StrPtr("IT")))

		})
	})
})
