package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("factorial:", factorial(10))
}

func factorial(n int32) string {
	r := []uint8{1}

	for i := int32(1); i <= n; i++ {
		r = multiplyNumber(r, uint8(i))
	}

	return toString(r)
}

func multiplyNumber(ar []uint8, d uint8) []uint8 {
	digits := digits(d)

	// multiply each digit separately
	var rows [][]uint8

	for i := 0; i < len(digits); i++ {
		rows = append(rows, multiplyDigit(ar, digits[i]))
	}

	// sum results with shift to left
	l := len(ar) + len(digits) + 1

	var r []uint8
	var h = uint8(0)
	for i := 0; i < l; i++ {

		sum := uint8(0)

		for j := 0; j < len(rows); j++ {
			ind := len(rows[j]) - 1 - i + j

			v := uint8(0)

			if 0 <= ind && ind < len(rows[j]) {
				v = rows[j][ind]
			}

			sum += v
		}

		sum = sum + h
		r = append(r, sum%10)
		h = sum / 10
	}

	return reverse(r)
}

func multiplyDigit(ar []uint8, d uint8) []uint8 {
	c := make([]uint8, len(ar)+1)

	overflow := uint8(0)

	for i := len(ar) - 1; i >= 0; i-- {
		v := d * ar[i]
		lowest := v % 10
		current := lowest + overflow
		c[i+1] = current % 10
		overflow = v/10 + current/10
	}

	if overflow > 0 {
		c[0] = overflow
	}

	return c
}

func reverse(r []uint8) []uint8 {
	var res []uint8
	nonZeroFound := false

	for i := len(r) - 1; i >= 0; i-- {
		if r[i] > 0 {
			nonZeroFound = true
		}

		if nonZeroFound {
			res = append(res, r[i])
		}
	}

	return res
}

func digits(d uint8) []uint8 {
	var r []uint8

	for {
		if d > 0 {
			r = append(r, d%10)
			d = d / 10
		} else {
			break
		}
	}

	return r
}

func toString(r []uint8) string {
	s := ""

	for i := 0; i < len(r); i++ {
		s += strconv.Itoa(int(r[i]))
	}

	return s
}
