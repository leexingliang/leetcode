/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head // 添加一个哨兵节点 (头结点)

	node := newHead
	for node != nil {
		if node.Next != nil && node.Next.Val == val {
			if node.Next.Next == nil {
				// 尾节点特殊处理
				node.Next = nil
				break
			} else {
				node.Next = node.Next.Next
			}
			// 当有两个连续的值, 如果不是 continue, 则会出错
			continue
		}
		node = node.Next
	}
	return newHead.Next
}