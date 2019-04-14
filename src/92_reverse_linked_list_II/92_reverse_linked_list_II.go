package reverse_linked_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	if m < 0 || n < 0 {
		return head
	}
	if m >= n {
		return head
	}
	if m < 0 || n < 0 {
		return head
	}
	if m >= n {
		return head
	}

	var index int // index

	// reverse list head
	reverse := new(ListNode)
	// prev node
	newHead := new(ListNode)
	newHead.Next = head
	prev := newHead

	for index = 1; index < m; index++ {
		if prev == nil {
			// list length is not enough
			return head
		}
		prev = prev.Next
	}

	// begin
	node := prev.Next
	if node == nil {
		return head
	}
	for index = m; index <= n; index++ {
		var tmp = node.Next

		node.Next = reverse.Next
		reverse.Next = node
		node = tmp
		if node == nil {
			break
		}
	}
	prev.Next.Next = node
	prev.Next = reverse.Next
	return newHead.Next
}
