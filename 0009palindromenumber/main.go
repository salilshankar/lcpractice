package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))

}

func isPalindrome(x int) bool {
    
    if x < 0 {
        return false
    }
    
    if x < 10 {
        return true
    }
    
    if x < 100 && x % 11 == 0 {
            return true
    }
    
    
    if x % 10 == 0 {
            return false
    }
    
    y := 0
    
    for x > 0 {
        y = (y*10) + (x%10)
        if x == y {
            return true
        }
        x /= 10
        
        if (x==y) {
            return true
        }
    }
    
    return false
    
}