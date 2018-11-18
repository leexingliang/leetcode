package remove_duplicate_nodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newHead *ListNode = new(ListNode)
	newHead.Next = head

	var node *ListNode = head    // 节点判断
	var tail *ListNode = newHead // 新链表结尾, 初始值为 newHead

	for node != nil {
		for node.Next != nil && node.Val == node.Next.Val {
			node = node.Next
		}
		if tail.Next == node {
			tail = tail.Next
		} else {
			tail.Next = node.Next
		}
		node = node.Next
	}

	return newHead.Next
}
