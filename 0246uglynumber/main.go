package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthUglyNumber(15))
}

func nthUglyNumber(n int) int {
    if n<=0 {
        return 0
    }
    
    ptr2, ptr3, ptr5 := 0, 0, 0
    
    x := make([]int, n)
    x[0] = 1
    
    for i:=1; i<n; i++ {
        x[i] = min(x[ptr2]*2,min (x[ptr3]*3, x[ptr5]*5))
        
        if x[i]==x[ptr2]*2 {
             ptr2++
        }
        
        if x[i]==x[ptr3]*3 {
            ptr3++
        }
        
        if x[i]==x[ptr5]*5 {
            ptr5++
        }
    }
    
    return (x[len(x)-1])
}

func min(a, b int) int {
    if a < b {
       return a
    }
    
    return b
}