package retry_test

import (
	"fmt"

	. "github.com/cthulhu/go-steun/retry"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Retry", func() {
	var operation *RetryOperation
	var Retried, AfterRetry bool
	var RetriedWithError, AfterRetryError error
	BeforeEach(func() {
		RetriedWithError = nil
		Retried = false
		AfterRetry = false
		AfterRetryError = nil
		operation = New()
	})
	Context("No error", func() {
		It("executes only main function", func() {
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
		It("executes main function and before retry", func() {
			operation.BeforeRetry(func(err error) {
				Retried = true
				RetriedWithError = err
			})
			operation.AfterRetry(func(err error) {
				AfterRetry = true
				AfterRetryError = err
			})
			operation.Do(func() error {
				return fmt.Errorf("Generic Error")
			})
			Expect(RetriedWithError).To(HaveOccurred())
			Expect(Retried).To(BeTrue())
			Expect(AfterRetry).To(BeTrue())
			Expect(AfterRetryError).To(HaveOccurred())
		})
	})
	Context("With error no retry", func() {
		It("executes main and before function", func() {
			Retried = false
			RetriedWithError = nil
			operation.Do(func() error {
				return fmt.Errorf("Generic Error")
			})
			Expect(RetriedWithError).NotTo(HaveOccurred())
			Expect(Retried).To(BeFalse())
		})
	})
})
