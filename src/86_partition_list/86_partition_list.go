package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var newLess *ListNode = new(ListNode)
	var less *ListNode = newLess

	var newBig *ListNode = new(ListNode)
	var big *ListNode = newBig

	var node *ListNode = head

	for node != nil {
		if node.Val < x {
			less.Next = node
			less = less.Next
		} else {
			big.Next = node
			big = big.Next
		}
		node = node.Next
	}
	less.Next = newBig.Next
	big.Next = nil

	return newLess.Next
}
