package linked_list_cycle_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 1. 计算环的长度
	cycleLength := cycleLen(hasCycle(head))
	if cycleLength == -1 {
		return nil
	}

	// 2. 找到环的入口节点
	fast, slow := head, head
	for ; cycleLength > 1; cycleLength-- {
		fast = fast.Next
	}
	for fast.Next != slow {
		fast = fast.Next
		slow = slow.Next
	}

	// 3. 去掉环
	// fast.Next = nil
	return slow
}

func hasCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head.Next, head

	for fast != nil && fast.Next != nil && slow != nil {
		if fast == slow {
			return fast
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}

func cycleLen(node *ListNode) int {
	if node == nil {
		return -1
	}

	stay := node
	move := stay.Next
	length := 1
	for stay != move {
		length++
		move = move.Next
	}
	return length
}
