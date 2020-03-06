package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Split Strings", func() {
	It("Basic tests", func() {
		Expect(Solution("abc")).To(Equal([]string{"ab", "c_"}))
		Expect(Solution("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
		Expect(Solution("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
	})
})

// 1. ginkgo bootstrap will gen suite_test
// 2. run go test or package from IDE directly
func Solution(str string) []string {
	if len(str)%2 != 0 {
		str += "_"
	}

	var l []string

	for i := 0; i < len(str); i += 2 {
		l = append(l, str[i:i+2])
	}

	return l
}
