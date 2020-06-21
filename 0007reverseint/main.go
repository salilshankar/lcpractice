package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(1534236469)
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {

	rem := 0
	rev := 0

	for x != 0 {
		rem = x % 10
		if rev > math.MaxInt32/10 {
			return 0
		} else if rev < math.MinInt32/10 {
			return 0
		}
		rev = rev*10 + rem
		x = x / 10
	}
	return rev

}
