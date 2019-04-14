package Swap_Nodes_in_Pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := new(ListNode)

	prev := newHead
	first := head
	second := head.Next

	for first != nil && second != nil {
		// 记录下一个 pair
		var tmp = second.Next

		prev.Next = second
		second.Next = first

		prev = first
		first = tmp
		prev.Next = first
		if first == nil {
			break
		}
		second = first.Next
	}
	return newHead.Next
}
