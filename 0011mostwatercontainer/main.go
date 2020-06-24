package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	area := 0
	ptr1 := 0
	ptr2 := len(height) - 1
	minheight := 0

	for i := 0; i < len(height); i++ {
		minheight = min(height[ptr1], height[ptr2])
		area = max(area, minheight*(ptr2-ptr1))

		if minheight == height[ptr1] {
			ptr1++
		} else {
			ptr2--
		}
	}

	return area

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
