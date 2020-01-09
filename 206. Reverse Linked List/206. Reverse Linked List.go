/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{} // 新链表头结点

	node := head
	for node != nil {
		tmp := node
		node = node.Next

		tmp.Next = newHead.Next
		newHead.Next = tmp
	}
	return newHead.Next
}