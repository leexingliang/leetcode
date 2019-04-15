package reverse_linked_list_II

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	if m < 0 || n < 0 {
		return head
	}
	if m >= n {
		return head
	}

	var prev *ListNode
	node := head
	for i := 1; i < m; i++ {
		prev = node
		node = node.Next
	}

	begin := prev
	tail := node
	for i := m; i <= n; i++ {
		var tmp = node.Next

		node.Next = prev
		prev = node
		node = tmp
	}
	if begin != nil {
		begin.Next = prev
	} else {
		head = prev
	}
	tail.Next = node
	return head
}
