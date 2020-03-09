package EqualizeArray

func equalizeArray(arr []int32) int32 {
	m := make(map[int32]int32)
	max := int32(0)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
		if m[arr[i]] > max {
			max = m[arr[i]]
		}
	}

	return int32(len(arr)) - max
}
