package structs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStructs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Structs Suite")
}
