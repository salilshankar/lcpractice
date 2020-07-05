package main

import (
	"fmt"
)

func main() {
	ans := generateParenthesis(4)
	fmt.Println(ans, "answer", len(ans))
}

func generateParenthesis(n int) []string {
	ans := make([][]string, 0)
	ans = append(ans, []string{""})
    if n == 0 {
        return ans[len(ans)-1]
    }
	
	i:=1
    for i<=n {
		ans=expand(ans)
		i++
	}

	for _, val := range ans {
		fmt.Println(val, "in ans")
	}
	//reversing is optional
    return reverse(ans[len(ans)-1])
    
}

func expand(ans [][]string) [][]string {
    x := []string{}
    for i := range ans {
		first := i
		second := len(ans) - i - 1
		if second == 0 {
			for _, str := range ans[i] {
				x = append(x, "(" + str + ")")
			}
		} else {
			x = append(x, getStrings(ans[first], ans[second])...)
		}
		
	}
	return append(ans, x)
}

func getStrings(a, b []string) []string {
	x := []string{}
		for _, str1 := range a {
			for _, str2 := range b {
				x = append(x, "(" + str1 + ")" + str2)
			}
		}

	return x

}

//optional. only if you want answer to be an exact match of the expectation. Otherwise, LC accepts non-reversed array too.
func reverse(a []string) []string {
	temp1:=""
	temp2:=""
	for i:=0; i<(len(a)-1)/2; i+=2 {
		temp1 = a[i]
		temp2 = a[i+1]
		a[i], a[i+1] = a[len(a)-i-1], a[len(a)-i-2]
		a[len(a)-i-1], a[len(a)-i-2] = temp1, temp2
	}
	return a
}

