package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "a*"))
}

func isMatch(s string, p string) bool {

	x := make([][]bool, len(s)+1)
	for i := range x {
		x[i] = make([]bool, len(p)+1)
	}

	x[0][0] = true

	for i := 1; i < len(x[0]); i++ {
		if p[i-1] == '*' {
			x[0][i] = x[0][i-2]
		}
	}

	for i := 1; i < len(x); i++ {
		for j := 1; j < len(x[0]); j++ {
			if (p[j-1] == '.') || (p[j-1] == s[i-1]) {
				x[i][j] = x[i-1][j-1]
			} else if p[j-1] == '*' {
				x[i][j] = x[i][j-2]
				if p[j-2] == '.' || (p[j-2] == s[i-1]) {
					x[i][j] = x[i][j] || x[i-1][j]
				}
			} else {
				x[i][j] = false
			}
		}
	}

	return x[len(s)][len(p)]

}
