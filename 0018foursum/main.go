package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(fourSum([]int{-3,-2,-1,0,0,1,2,3}, 0))

}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	m := make(map[[4]int]int)
	ctr := 0
	for i := 0; i < len(nums)-3; i++ {
		j := i+1
		for j <= len(nums) - 3  {
			k := j+1 
			l := len(nums)-1
			fmt.Println()
			for k < l {
				fmt.Println("indices:", i,j,k,l)
				fmt.Println("values:", nums[i],nums[j],nums[k],nums[l])
				a, b, y, z := nums[i], nums[j], nums[k], nums[l]
				if a+b+y+z == target {
					if _, ok := m[[4]int{a,b,y,z}]; !ok {
						result = append(result, []int{a, b, y, z})
						ctr++
						m[[4]int{a,b,y,z}] = ctr
					}
					l--
					k++
				} else if a+b+y+z > target {
					l--
				} else if a+b+y+z < target {
					k++
				}
			}
			j++
		}	
	}
	return result
}