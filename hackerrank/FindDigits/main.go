package FindDigits

import "strconv"

func findDigits(n int32) int32 {
	l := len(strconv.Itoa(int(n)))
	c := int32(0)
	rest := n
	var last int32
	for i := 0; i < l; i++ {
		last, rest = rest%10, rest/10

		if last != 0 && n%last == 0 {
			c++
		}
	}

	return c
}

func findDigits2(n int32) int32 {
	c := int32(0)
	rest := n
	var last int32
	for {
		if rest <= 0 {
			break
		}

		last, rest = rest%10, rest/10

		if last != 0 && n%last == 0 {
			c++
		}
	}

	return c
}
