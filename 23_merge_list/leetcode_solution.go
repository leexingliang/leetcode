package merge_list

// 采用递归分治思想
func mergeKLists_leetcode(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	return merge(mergeKLists_leetcode(lists[:length/2]), mergeKLists_leetcode(lists[length/2:]))
}

func merge(listA, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	}
	if listB == nil {
		return listA
	}
	head := new(ListNode)
	node := head
	for listA != nil && listB != nil {
		if listA.Val < listB.Val {
			node.Next = listA
			listA = listA.Next
		} else {
			node.Next = listB
			listB = listB.Next
		}
		node = node.Next
	}
	if listA != nil {
		node.Next = listA
	}
	if listB != nil {
		node.Next = listB
	}
	return head.Next
}
