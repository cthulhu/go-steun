package pointers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPointers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pointers Suite")
}
