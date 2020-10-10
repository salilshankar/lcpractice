package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("10", "110010"))
}

func addBinary(a string, b string) string {

	if a == "0" && a == b {
		return "0"
	}

	if a == "1" && a == b {
		return "10"
	}

	if len(a) == len(b) && len(a) == 1 {
		return "1"
	}

	lenA := len(a)
	lenB := len(b)

	ans := []byte{}

	if lenA > lenB {
		ans = addProper(a, b)
	} else {
		ans = addProper(b, a)
	}

	return string(ans)
}

func add(a, b, carry byte) []byte {
	str := make([]byte, 2)
	str[0] = 48
	str[1] = 49
	if a == b && a == str[0] {
		return addCarry([]byte{48, 48}, carry)
	} else if a == b && a == str[1] {
		return addCarry([]byte{49, 48}, carry)
	}

	return addCarry([]byte{48, 49}, carry)
}

func addCarry(a []byte, carry byte) []byte {
	if carry == 48 {
		return a
	}

	if a[0] == a[1] && a[0] == 48 {
		return []byte{48, 49}
	}

	if a[1] == 48 {
		return []byte{49, 49}
	}

	return []byte{49, 48}

}

func addProper(larger, smaller string) []byte {
	smallerMatched := make([]byte, len(larger))
	for i := range smaller {
		smallerMatched[len(smallerMatched)-1-i] = smaller[len(smaller)-1-i]
	}

	for i := range smallerMatched {
		if i < len(smallerMatched)-len(smaller) {
			smallerMatched[i] = 48
		}
	}

	ans := make([]byte, len(larger)+1)
	ans[0] = 48
	var carry byte
	carry = 48
	for i := range larger {
		sum := add(larger[len(larger)-1-i], smallerMatched[len(smallerMatched)-1-i], carry)
		carry = sum[0]
		ans[len(ans)-1-i] = sum[1]
	}

	ans[0] = carry

	if ans[0] == 48 {
		ans = ans[1:]
	}

	return ans
}
