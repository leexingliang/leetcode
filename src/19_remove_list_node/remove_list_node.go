package remove_list_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n < 0 {
		return head
	}

	var tail *ListNode = head
	var node *ListNode = head

	var i int
	for i = 1; i <= n; i++ {
		if tail.Next != nil {
			tail = tail.Next
		}
	}
	if i < n {
		return head
	}

	// 遍历,找到倒数第 n+1 个节点
	for tail.Next != nil {
		node = node.Next
		tail = tail.Next
	}

	node.Next = node.Next.Next

	return head
}
