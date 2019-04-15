// 反转链表
// 每 k 个节点做一次反转, 如果节点数小于 K 个, 不做处理
// 假设有链表节点为 [1   2   3   4   5], k 为3, 反转之后
// 3   2   1   4   5
package reverse_list_kNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路:
// 借助一个额外的 node 节点, 每次遍历 k 个节点, 反转之后加入到新的链表中
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k <= 1 {
		return head
	}

	var newHead *ListNode = new(ListNode)
	node := newHead

	var slow *ListNode = head
	var fast *ListNode = head
	index := 1

	for slow != nil {
		for index = 1; index <= k; index++ {
			slow = slow.Next
			if slow == nil && index < k {
				goto end
			}
		}
		// 反转 k 个节点
		for fast != slow {
			tmp := fast
			fast = fast.Next
			tmp.Next = node.Next
			node.Next = tmp
		}
		// 定位到新链表结尾
		for node.Next != nil {
			node = node.Next
		}
	}
end:
	if slow == nil && index < k {
		node.Next = fast
	}

	return newHead.Next
}
