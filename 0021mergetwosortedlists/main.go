package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	l1ptr := new(ListNode)
	l1ptr = l1
	l1.Val = 2
	l1.Next = new(ListNode)
	l1 = l1.Next
	l1.Val = 3
	l1.Next = new(ListNode)
	l1 = l1.Next
	l1.Val = 4
	l1.Next = nil

	l2 := new(ListNode)
	l2ptr := new(ListNode)
	l2ptr = l2
	l2.Val = 1
	l2.Next = new(ListNode)
	l2 = l2.Next
	l2.Val = 3
	l2.Next = new(ListNode)
	l2 = l2.Next
	l2.Val = 5
	l2.Next = nil

	ans := mergeTwoLists(l1ptr, l2ptr)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {
        return nil
    }
    
    if l1 == nil {
        return l2
    }
    
    if l2 == nil {
        return l1
    }
    
    lpointer := new(ListNode)
    temp := new(ListNode)
    
    if l1.Val > l2.Val {
        temp = l1
        l1 = l2
        l2 = temp
    }

	lpointer = l1
    
    for l1 != nil && l2 != nil {
        for l1.Next != nil && l1.Next.Val <= l2.Val {
            l1 = l1.Next
        }
        temp = l1.Next
        l1.Next = l2
        l2 = temp
    }
    
    
    return lpointer
    
}

