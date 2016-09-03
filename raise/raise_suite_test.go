package raise_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestErr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Raise Suite")
}
