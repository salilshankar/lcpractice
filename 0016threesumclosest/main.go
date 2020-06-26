package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100))
}

func threeSumClosest(nums []int, target int) int {
	distance := math.MaxInt64
	sort.Ints(nums)
	maxIndex := len(nums) - 1
	low := 0
	high := 0
	for i := 0; i < len(nums); i++ {
		low = i + 1
		high = maxIndex
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if abs(distance) > abs(target-sum) {
				distance = target - sum
			}
			if sum < target {
				low++
			} else {
				high--
			}

		}
	}
	return target - distance
}

func abs(a int) int {
	if a < 0 {
		return (-1 * a)
	}
	return a
}
