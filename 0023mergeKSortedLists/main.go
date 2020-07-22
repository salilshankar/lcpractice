package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

	list1 := ListNode{
		3, 
		new(ListNode),
	}

	list1.Next = &ListNode{
		3, 
		new(ListNode),
	}

	list1.Next.Next = &ListNode{
		4, 
		nil,
	}

	list2 := ListNode{
		4, 
		nil,
	}

	list2.Next = &ListNode{
		5, 
		new(ListNode),
	}

	list2.Next.Next = &ListNode{
		10, 
		nil,
	}

	ans := mergeKLists([]*ListNode{&list1, &list2})

	for ans.Next != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

	fmt.Println(ans.Val)


}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    
    l1 := lists[0]
    for i:=1; i<len(lists); i++ {
        if l1 != nil && lists[i] != nil {
            l1 = mergeLists(l1, lists[i])
        } else if l1 == nil && lists[i] != nil {
            l1 = lists[i]
        }
        
    }
    
    return l1
    
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	ptr := new(ListNode)
	
	if list1.Val > list2.Val {
        temp := list1
		list1 = list2
		list2 = temp
	}

    ans := list1
    
    for list1 != nil && list2 != nil {
		for list1.Next != nil && list1.Next.Val <= list2.Val {
			list1 = list1.Next
		}
		
		ptr = list2.Next
		list2.Next = list1.Next
		list1.Next = list2
		list2 = ptr
    }
    
    return ans
}