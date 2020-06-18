package main

import (
	"fmt"
	"math"
)

func main() {

	a1 := []int{1, 2}
	a2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(a1, a2))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var median float64
	var sx, lx, sy, ly float64
	if len(nums1) <= len(nums2) {
		sx, lx, sy, ly = findMedianCenter (nums1, nums2, 0, len(nums1))
	} else {
		sx, lx, sy, ly = findMedianCenter (nums2, nums1, 0, len(nums2))
	}


	if (len(nums1)+len(nums2)) % 2 == 0 {
		median = average(max(sx, sy), min(lx, ly))
	} else {
		median = max(sx, sy)
	}
	
	return median
}

func findMedianCenter(x []int, y []int, start int, end int) (float64, float64, float64, float64){
	partx := (start + end) / 2
	party := (len(x)+len(y)+1)/2 - partx

	var x1, x2, y1, y2 float64

	if partx-1 >= 0 {
		x1 = float64(x[partx-1])
	} else {
		x1 = math.Inf(-1)
	}

	if partx < len(x) {
		x2 = float64(x[partx])
	} else if partx == len(x) {
		x2 = math.Inf(1)
	}

	if party-1 >= 0 {
		y1 = float64(y[party-1])
	} else {
		y1 = math.Inf(-1)
	}

	if party < len(y) {
		y2 = float64(y[party])
	} else if party == len(y) {
		y2 = math.Inf(5)
	}

	shouldMove := isMedianCenter(x1, x2, y1, y2)
	

	switch shouldMove {
	case 0:
		return x1, x2, y1, y2
	case 1:
		x1, x2, y1, y2 = findMedianCenter(x, y, start+1, end) 
	case -1:
		x1, x2, y1, y2 = findMedianCenter(x, y, start, end-1)
	}

	return x1, x2, y1, y2

}

func isMedianCenter(x1, x2, y1, y2 float64) int {

	if (x1 <= y2) && (y1 <= x2) {
		return 0
	} else if x1 > y2 {
		return -1
	}
	return 1
}

func average(a float64, b float64) float64 {
	return (a + b) / float64(2)
}

func max(a float64, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

func min(a float64, b float64) float64 {
	if a < b {
		return a
	}

	return b
}
