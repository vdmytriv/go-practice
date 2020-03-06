package kata_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreatePhoneNumber()", func() {
	It("basic test", func() {
		Expect(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})).To(Equal("(123) 456-7890"))
	})
})

func CreatePhoneNumber(numbers [10]uint) string {
	parts := [3]string{}
	index := 0

	for i := 0; i < 10; i++ {
		if i >= 3 {
			index = 1
		}
		if i >= 6 {
			index = 2
		}

		parts[index] = parts[index] + fmt.Sprint(numbers[i])
	}

	return fmt.Sprintf("(%s) %s-%s", parts[0], parts[1], parts[2])
}
