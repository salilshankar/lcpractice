package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{3,4,5}
	arr2 := []int{2, 7}

	fmt.Println(findMedianSortedArrays(arr1,arr2))
}





func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var median float64
    nums := append(nums1, nums2...)
    sort.Ints(nums)
    
    if len(nums) % 2 == 0 {
        halflen := len(nums)/2
        median = (float64(nums[halflen]) + float64(nums[halflen-1]))/2
    } else {
        median = float64(nums[(len(nums)-1)/2])
    }
    
    return median
}