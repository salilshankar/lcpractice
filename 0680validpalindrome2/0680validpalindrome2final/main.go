package main

import (
	"fmt"
)

func main() {
	fmt.Println(validPalindrome("eccer"))
}

func validPalindrome(s string) bool {

	lenStr := len(s)

	if lenStr == 1 {
		return true
	}

	if lenStr == 2 {
		return true
	}

	isPalindrome, failIndex := checkPalindrome(s)

	if isPalindrome {
		return true
	}

	str1 := string(s[:failIndex] + s[failIndex+1:])
	str2 := string(s[:lenStr-1-failIndex] + s[lenStr-failIndex:])
	
	p1, _ := checkPalindrome(str1)
	p2, _ := checkPalindrome(str2)

	return p1 || p2
}

func checkPalindrome(s string) (bool, int) {

	halfLen := len(s) / 2

	for i := 0; i < halfLen; i++ {
		if s[i] != s[len(s)-1-i] {
			return false, i
		}
	}

	return true, halfLen
}
