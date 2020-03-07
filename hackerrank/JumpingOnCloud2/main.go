package JumpingOnCloud2

func jumpingOnClouds(c []int32, k int32) int32 {
	res := int32(100)

	i := k % int32(len(c))
	res -= c[i]*2 + 1

	for i != 0 {
		i = (i + k) % int32(len(c))
		res -= c[i]*2 + 1
	}

	return res
}
