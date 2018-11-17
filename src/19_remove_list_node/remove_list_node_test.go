package remove_list_node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createList(num int) *ListNode {
	var head *ListNode = new(ListNode)
	tail := head

	for i := 1; i <= num; i++ {
		node := new(ListNode)
		node.Val = i
		node.Next = nil

		tail.Next = node
		tail = node
	}

	return head.Next
}

func getListVal(head *ListNode) []int {
	result := make([]int, 0)

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func Test_RemoveNode(t *testing.T) {
	list := createList(5)

	removeNthFromEnd(list, 2)

	assert.Equal(t, []int{1, 2, 3, 5}, getListVal(list))
}
