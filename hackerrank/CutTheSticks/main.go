package CutTheSticks

import "sort"

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int32) []int32 {
	m := make(map[int]int)
	l := []int{}
	for i := 0; i < len(arr); i++ {
		m[int(arr[i])]++
		if m[int(arr[i])] == 1 {
			l = append(l, int(arr[i]))
		}
	}

	sort.Ints(l)

	sum := int32(len(arr))
	res := []int32{sum}
	for i := 0; i < len(l)-1; i++ {
		sum = sum - int32(m[l[i]])
		res = append(res, sum)
	}

	return res
}
