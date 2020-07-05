func hammingDistance(x int, y int) int {
	temp := x ^ y
	ctr := 0
	for temp > 0 {
		if (temp & 1) > 0 {
			ctr++
		}
		temp >>= 1
	}

	return ctr
}