func bitwiseComplement(N int) int {

	if N == 0 {
		return 1
	}

	bits := int(math.Log2(float64(N)))

	for i := 0; i <= bits; i++ {
		N = N ^ (1 << i)
	}

	return N
}