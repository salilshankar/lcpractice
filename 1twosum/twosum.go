package main

import (
	"fmt"
)

func main() {
	testarr := []int{3, 2, 3}

	fmt.Println(twoSum(testarr, 6))

}

func twoSum(nums []int, target int) []int {

	diff := make(map[int]int)

	for i, v := range nums {
		if ans, ok := diff[target-v]; ok {
			return []int{ans, i}
		}

		diff[v] = i
	}

	return nil

}
