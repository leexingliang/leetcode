package reverse_linked_list_II

import (
	"fmt"
	"testing"
)

func createList(start, end int) *ListNode {
	var head *ListNode = new(ListNode)
	tail := head

	for i := start; i <= end; i++ {
		node := new(ListNode)
		node.Val = i
		node.Next = nil

		tail.Next = node
		tail = node
	}

	return head.Next
}

func Test_reverse_1(t *testing.T) {
	var head *ListNode = new(ListNode)

	node := new(ListNode)
	node.Val = 5
	node.Next = head.Next
	head.Next = node

	node = new(ListNode)
	node.Val = 3
	node.Next = head.Next
	head.Next = node

	ret := reverseBetween2(head.Next, 1, 2)
	fmt.Println(ret.Val, ret.Next.Val)
}
