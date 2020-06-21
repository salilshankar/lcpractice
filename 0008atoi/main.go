package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("  -0012a42"))
}

func myAtoi(str string) int {
	sign := true
	var num int
	numbercalc := false

	for _, char := range str {

		if numbercalc && !(char > 47 && char < 58) {
			if !sign {
				return (0 - num)
			}
			return num
		}
		if !((char > 47 && char < 58) || (char == 43 || char == 45) || (char == 32)) {
			if !sign {
				return (0 - num)
			}
			return num
		}

		if char == 43 {
			sign = true
			numbercalc = true
		} else if char == 45 {
			sign = false
			numbercalc = true
		} else if char > 47 && char < 58 {
			if (num > math.MaxInt32/10) && sign {
				return math.MaxInt32
			} else if (0-num < math.MinInt32/10) && !sign {
				return math.MinInt32
			}
			numbercalc = true
			num = num*10 + int((char - 48))
		}

	}

	if (num > math.MaxInt32) && sign {
		return math.MaxInt32
	} else if (0-num < math.MinInt32) && !sign {
		return math.MinInt32
	}

	if !sign {
		return (0 - num)
	}

	return num

}
