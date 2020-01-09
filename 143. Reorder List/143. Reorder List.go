/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {

	var frontHalfNode Stack // 链表前半部分节点

	var fast *ListNode
	var slow *ListNode
	var behindHalfNode *ListNode
	// var length int32 // 链表长度, 判断奇偶使用

	newHead := &ListNode{} // 新链表头结点

	// 利用快慢指针方式找到中间节点 (区分奇偶数)
	//
	fast, slow = head, head
	for fast != nil && fast.Next != nil {
		frontHalfNode.Push(slow) // 前半部分节点入栈

		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		// 奇数个节点
		behindHalfNode = slow.Next
		slow.Next = nil
		newHead.Next = slow
	} else {
		// 偶数个节点
		behindHalfNode = slow
	}

	for {
		nodeFront := frontHalfNode.Pop()
		if nodeFront == nil {
			// 没了
			break
		}
		nodeBehind := behindHalfNode
		if nodeBehind == nil {
			break
		}
		behindHalfNode = behindHalfNode.Next

		nodeFront.Next = nodeBehind
		nodeBehind.Next = newHead.Next
		newHead.Next = nodeFront
	}
	head = newHead.Next
}

type Stack struct {
	data []*ListNode

	len int32
	cap int32
}

func (s *Stack) Push(node *ListNode) {
	if s.len == s.cap {
		s.data = append(s.data, node)
		s.len++
		s.cap++
	} else if s.len < s.cap {
		s.data[s.len] = node
		s.len++
	}
}

func (s *Stack) Pop() *ListNode {
	if s.len == 0 {
		return nil
	}

	s.len--
	return s.data[s.len]
}