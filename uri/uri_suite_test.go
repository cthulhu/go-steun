package uri_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUri(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uri Suite")
}
