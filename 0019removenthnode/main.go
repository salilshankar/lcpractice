/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	lpointer := new(ListNode)
	first, second := new(ListNode), new(ListNode)

	lpointer.Next = head
	first, second = lpointer, lpointer

	for i := 1; i <= n; i++ {
		first = first.Next
	}

	first = first.Next

	for first != nil {
		second = second.Next
		first = first.Next
	}

	second.Next = second.Next.Next

	return lpointer.Next

}