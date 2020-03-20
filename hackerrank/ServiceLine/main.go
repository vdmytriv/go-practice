package main

// Complete the serviceLane function below.
func serviceLane(n int32, width []int32, cases [][]int32) []int32 {
	var res []int32

	for i := 0; i < len(cases); i++ {

		min := width[cases[i][0]]
		for j := cases[i][0]; j <= cases[i][1]; j++ {
			if width[j] < min {
				min = width[j]
			}
		}

		res = append(res, min)
	}

	return res
}
