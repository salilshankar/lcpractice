package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))

}

func convert(s string, numRows int) string {
    var ctr int
    str := make([]string, numRows)
    direction := -1
    
    for _, c := range s {
		fmt.Println(ctr)
        str[ctr] = str[ctr]+string(c)
       
        if ctr == 0 {
            direction *= -1
		} 
		if (ctr+1) % numRows == 0 {
            direction *= -1
		}
		
		ctr += direction
		
		
    }
    
    retstr := ""
    
    for _, instr := range str {
		fmt.Println(instr)
        retstr = retstr + instr
    }
    
    return retstr
}
