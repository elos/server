package util_test

import (
	. "github.com/elos/server/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Random", func() {
	Context("RandomString()", func() {
		It("should return random strings", func() {
			var (
				testNeg1 string = RandomString(-1)
				test0    string = RandomString(0)
				test26   string = RandomString(26)
				test256  string = RandomString(256)
			)

			Expect(testNeg1).To(HaveLen(0))
			Expect(test0).To(HaveLen(0))
			Expect(test26).To(HaveLen(26))
			Expect(test256).To(HaveLen(256))
		})
	})
})
