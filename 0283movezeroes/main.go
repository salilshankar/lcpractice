func moveZeroes(nums []int) {
	index := 0

	for _, val := range nums {
		if val != 0 {
			nums[index] = val
			index++
		}
	}

	for index < len(nums) {
		nums[index] = 0
		index++
	}

}

