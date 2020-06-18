package main

import (
	"fmt"
)

func main() {
	fmt.Println("222020221")
	fmt.Println()
	fmt.Println()
	fmt.Println(longestPalindrome("222020221"))
}

func longestPalindrome(s string) string {
	slength := len(s)

	if slength == 0 {
		return ""
	} else if slength == 1 {
		return s
	} else if slength == 2 {
		if s[0] == s[1] {
			return s
		}
		return string(s[0])
	} else if slength == 3 {
		if s[0] == s[2] {
			return s
		}
	} else if slength%2 == 0 {
		if s[:slength/2] == s[slength/2:] {
			return s
		}
	}
	lp := s[:1]
	tlp := ""
	
	for center := 1; center <= slength; center++ {
		if tlp = findPalindrome(s, center, center); len(lp) < len(tlp) {
			lp = tlp
		}

		if tlp = findPalindrome(s, center, center+1); len(lp) < len(tlp) {
			lp = tlp
		}


	}
	return lp
}

func findPalindrome(s string, left int, right int) string {
	templp := ""
	for left>=0 && right < len(s) && s[left]==s[right] {
		templp = s[left:right+1]
		left--
		right++	
	}

	return templp
}
