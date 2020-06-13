package main

import (
	"fmt"
)

func main() {

	longestSubstr := lengthOfLongestSubstring("abcabcbb")

	fmt.Println(longestSubstr)

}

func lengthOfLongestSubstring(s string) int {

	if len(s) == 1 {
		return 1
	}

	if len(s) == 2 && (s[0]) != s[1] {
		return 2
	} else if len(s) == 2 {
		return 1
	}

	if len(s) == 0 {
		return 0
	}

	
	tail := -1
	m := make(map[string]int)
	maxlength := 0

	for head := 0; head < len(s); head++ {
		if v, ok := m[s[head:head+1]]; ok {
			if tail < v {
				tail = v
			}
		}

		m[s[head:head+1]] = head

		if head - tail > maxlength {
			maxlength = head - tail 
		}

	}	

	return maxlength

}
