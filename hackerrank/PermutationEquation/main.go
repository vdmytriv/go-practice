package main

func permutationEquation(p []int32) []int32 {
	m := make(map[int32]int32)
	for i := int32(1); i <= int32(len(p)); i++ {
		m[p[i-1]] = i
	}

	var res []int32
	for i := 1; i <= len(p); i++ {
		res = append(res, m[m[m[p[i-1]]]])
	}

	return res
}
