package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"abbb", "a", "accc", "aa"}))
}

func longestCommonPrefix(strs []string) string {
    
    if len(strs) == 0 {
       return ""
    }
    
    if len(strs) == 1 {
        return strs[0]
    }
    
    maxlen := min(len(strs[0]), len(strs[1]))
    
    if len(strs)>=3 {
    
    for i:=2; i<len(strs); i++{
        maxlen = min(maxlen,len(strs[i]))
    }
    }
    
    var chars []byte
    
   for i:=0; i<maxlen; i++ {
		diffChar := false
        for j:=0; j<len(strs)-1; j++ {
            if strs[j] == "" || strs[j+1] == "" {
                diffChar = true
            } else if strs[j][i] != strs[j+1][i] {
                diffChar = true
            }
            
		}
		
		if !diffChar {
			chars = append(chars, strs[0][i])
		} else {
			return string(chars)
		}
    }
    return string(chars)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}