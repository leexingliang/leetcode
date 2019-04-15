package rotate_list

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}

	var newHead *ListNode = new(ListNode)
	newHead.Next = head
	var fast *ListNode = head
	var slow *ListNode = head
	var index int

begin:
	for index = 1; index <= k; index++ {
		fast = fast.Next
		if fast == nil {
			break
		}
	}
	if fast == nil {
		if index == k {
			return head
		} else {
			k = k % index
			fmt.Println(index)
			fast = head
			goto begin
		}
	}
	// 找到倒数第 k 个节点
	// fast 指向链表最后一个节点
	// slow 指向倒数第 k+1 个节点
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 截断原有的链表
	fast.Next = newHead.Next
	newHead.Next = slow.Next
	slow.Next = nil

	return newHead.Next
}
