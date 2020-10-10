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
		ans = addBinaryExt(a, b)
	} else {
		ans = addBinaryExt(b, a)
	}

	return string(ans)
}

func addBinaryExt(larger, smaller string) []byte {
	bytes := "01"
	zero := bytes[0]

	// "Make another string called smalleMatched, filled with zeros, whose length matches to the larger string.  And then, copy the contents of smaller one from the bottom. For example, larger string 1010 and smaller string 11. So the smallerMatched contains 0011."
	smallerMatched := make([]byte, len(larger))
	for i := range smaller {
		smallerMatched[len(smallerMatched)-1-i] = smaller[len(smaller)-1-i]
	}

	for i := range smallerMatched {
		if i < len(smallerMatched)-len(smaller) {
			smallerMatched[i] = zero
		}
	}

	//do a bitwise add factoring in carry
	ans := make([]byte, len(larger)+1)
	ans[0] = zero
	var carry byte
	carry = zero
	for i := range larger {
		sum := add(larger[len(larger)-1-i], smallerMatched[len(smallerMatched)-1-i], carry)
		carry = sum[0]
		ans[len(ans)-1-i] = sum[1]
	}

	ans[0] = carry

	if ans[0] == zero {
		ans = ans[1:]
	}

	return ans
}

func add(a, b, carry byte) []byte {

	bytes := "01"
	zero := bytes[0]
	one := bytes[1]

	if a == b && a == zero {
		// "0 + 0 + carry"
		return addCarry([]byte{zero, zero}, carry)
	} else if a == b && a == one {
		// "1 + 1 + carry"
		return addCarry([]byte{one, zero}, carry)
	}

	// "1 + 0 + carry or 0 + 1 + carry"
	return addCarry([]byte{zero, one}, carry)
}

func addCarry(a []byte, carry byte) []byte {

	bytes := "01"
	zero := bytes[0]
	one := bytes[1]

	// "carry 0 has no consequence"
	if carry == zero {
		return a
	}

	// "00 + 1"
	if a[0] == a[1] && a[0] == zero {
		return []byte{zero, one}
	}

	//  "10 + 1"
	if a[1] == zero {
		return []byte{one, one}
	}

	// "01 + 1"
	return []byte{one, zero}

}
