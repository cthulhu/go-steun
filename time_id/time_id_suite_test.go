package time_id_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTimeId(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TimeId Suite")
}
