package main

import (
	"fmt"
)

//ListNode is a struct
type ListNode struct {
	     Val int
	     Next *ListNode
 }

 func main() {

	i1 :=[]int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	
	i2 := []int{5,6,4}

	num1 := inputGen (i1)
	num2 := inputGen (i2)

	/* num1 := ListNode{
		 Val: 0,
	 }

	 num2 := ListNode {
		 Val: 1,
	 } */

	 fmt.Println(addTwoNumbers(num1, num2))
 }

 func inputGen(l1 []int) *ListNode {
	ilist := new (ListNode)

	listMem := ilist

	for i := len(l1) - 1; i>=0; i-- {
		ilist.Val = l1[i]
		ilist.Next = new(ListNode)
		ilist = ilist.Next
	}

	return listMem

 }



 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    num1 := retNum (l1)
	num2 := retNum (l2)
	
	var sum uint64
    
	sum = num1+num2
	fmt.Println(sum)
    
    return retList(sum)
    
    
}

func retNum (l *ListNode) uint64 {
    var operator uint64 = 1
    var sum uint64
    for l != nil {
        sum = uint64(l.Val)*operator + sum
        l=l.Next
        a := operator
		operator *= 10
		//fmt.Println(operator)
        if ((operator/a)!=10) {
			//handle multiplication overflow
			operator = 0
			for i := 0; i < 1000; i++ {
				operator += a/100
				//fmt.Println(operator)
			}
			fmt.Println ("overflow checked", operator)
        }
        
	}
	
	fmt.Println(sum)
    
    return sum
}

func retList (sum uint64) *ListNode {
    slist := new(ListNode)
    var rem uint64
    listMem := slist
    
    for sum >= 1 {
        rem = sum % 10
        sum = sum / 10
        

        slist.Val = int(rem)
        if sum !=0 {
            slist.Next = new (ListNode)
            slist = slist.Next 
        }
        
        
    }
    
    slist = nil
    
    return listMem
    
    
}