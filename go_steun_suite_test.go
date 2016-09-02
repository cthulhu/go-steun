package go_steun_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoSteun(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoSteun Suite")
}
