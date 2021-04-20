// +build linux

package osutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDistroInfo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Linux distro test suite")
}

var _ = Describe("Linux distro unit test",
	func() {
		It("one should return 1", func() {
			Expect(1).To(Equal(1))
		})
	},
)
