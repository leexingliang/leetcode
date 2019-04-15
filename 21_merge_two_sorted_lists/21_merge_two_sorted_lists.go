package merge_two_sorted_lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
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

	newHead := l1
	node := l2
	if l1.Val > l2.Val {
		newHead = l2
		node = l1
	}
	prev := newHead
	next := newHead.Next

	for node != nil && next != nil {
		if next.Val <= node.Val {
			prev = next
			next = next.Next
		} else {
			var tmp *ListNode
			tmp = node.Next

			node.Next = prev.Next
			prev.Next = node
			prev = node

			node = tmp
		}
	}
	if node != nil {
		prev.Next = node
	}
	return newHead
}
