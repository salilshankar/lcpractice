package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{()"))
}

func isValid(s string) bool {
    if s=="" {
        return true
    }
    
    if len(s) == 1 {
        return false
    }
    
    m := make(map[byte]byte)
    var notClosed []byte
    
    m['('] = ')'
    m['['] = ']'
    m['{'] = '}'
    
    if len(s) == 2 && s[1] != m[s[0]] {
        return false
	} else if len(s) == 2 && s[1] == m[s[0]] {
        return true
    }
	 
    i := 1
    notClosed = append(notClosed, m[s[0]])
    for i<len(s) {
        if len(notClosed) != 0 && s[i] == notClosed[len(notClosed)-1] {
			notClosed = notClosed[:len(notClosed)-1]
			i++
		} else if v, ok := m[s[i]]; ok{
			notClosed = append(notClosed, v)
			i++
		} else {
			return false
		}
 	}
	
	if len(notClosed) > 0 {
        return false
	}
	
    return true
    
}