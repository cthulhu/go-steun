package retry_test

import (
	"fmt"

	. "github.com/cthulhu/go-steun/retry"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Retry", func() {
	operation := New()
	var Retried bool
	var RetriedWithError error
	BeforeEach(func() {
		RetriedWithError = nil
		Retried = false
	})
	Context("No error", func() {
		It("excutes only main function", func() {
			operation.BeforeRetry(func(err error) {
				Retried = true
				RetriedWithError = err
			})
			operation.Do(func() error {
				return nil
			})
			Expect(RetriedWithError).To(BeNil())
			Expect(Retried).To(BeFalse())
		})
	})
	Context("With error", func() {
		It("excutes only main function", func() {
			operation.BeforeRetry(func(err error) {
				Retried = true
				RetriedWithError = err
			})
			operation.Do(func() error {
				return fmt.Errorf("Generic Error")
			})
			Expect(RetriedWithError).To(HaveOccurred())
			Expect(Retried).To(BeTrue())
		})
	})
	Context("With error no retry", func() {
		It("excutes only main function", func() {
			operation.Do(func() error {
				return fmt.Errorf("Generic Error")
			})
			Expect(RetriedWithError).To(HaveOccurred())
			Expect(Retried).To(BeTrue())
		})
	})
})
