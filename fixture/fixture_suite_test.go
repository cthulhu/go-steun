package fixture_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fixture Suite")
}
