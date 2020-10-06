package main

import (
	"fmt"
)

func main() {
	fmt.Println(validPalindrome("cbbcc"))
}

func validPalindrome(s string) bool {

	lenStr := len(s)

	if lenStr == 1 {
		return true
	}

	if lenStr == 2 {
		return true
	}

	isPalindrome, upperHalf, lowerHalf := checkPalindrome(s)

	if isPalindrome {
		return true
	}

	//fmt.Println(upperHalf, lowerHalf)
	misMatchIndex := []int{}

	for i := range upperHalf {
		if upperHalf[i] != lowerHalf[i] {
			misMatchIndex = append(misMatchIndex, i)
			break
		}
	}

	if len(misMatchIndex) == 0 {
		return false
	}

	newUpperHalf := make([]byte, len(upperHalf)-1)
	k := 0
	for i := range upperHalf {
		if i != misMatchIndex[0] {
			newUpperHalf[k] = upperHalf[i]
			k++
		}
	}

	str1 := make([]byte, len(s)-1)
	for i, char := range newUpperHalf {
		str1[i] = char
	}

	for i, char := range invertBottomHalf(lowerHalf) {
		str1[i+len(newUpperHalf)] = byte(char)
	}

	isPalindrome, _, _ = checkPalindrome(string(str1))
	if isPalindrome {
		return true
	}

	newlowerHalf := make([]byte, len(lowerHalf)-1)
	k = 0
	for i := range lowerHalf {
		if i != misMatchIndex[0] {
			newlowerHalf[k] = lowerHalf[i]
			k++
		}
	}

	str1 = make([]byte, len(s)-1)
	for i, char := range upperHalf {
		str1[i] = byte(char)
	}

	for i, char := range invertBottomHalf(string(newlowerHalf)) {
		str1[i+len(upperHalf)] = byte(char)
	}

	isPalindrome, _, _ = checkPalindrome(string(str1))
	if isPalindrome {
		return true
	}

	return false
}

func checkPalindrome(s string) (bool, string, string) {

	halfLen := len(s) / 2

	upperHalf := s[:halfLen]
	bottomHalf := invertBottomHalf(s[halfLen:])

	tempbottomHalf := bottomHalf

	if len(tempbottomHalf) > len(upperHalf) {
		tempbottomHalf = tempbottomHalf[:len(tempbottomHalf)-1]
	}

	if upperHalf == tempbottomHalf {
		return true, upperHalf, bottomHalf
	}

	return false, upperHalf, bottomHalf
}

func invertBottomHalf(s string) string {
	invertedStr := make([]byte, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		invertedStr[len(s)-1-i] = s[i]
	}

	return string(invertedStr)

}
