package remove_duplicate_nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createList(start, end int) *ListNode {
	var head *ListNode = new(ListNode)
	tail := head

	for i := start; i <= end; i++ {
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

func Test_DeleteDuplicates(t *testing.T) {
	list := createList(1, 5)
	ret := deleteDuplicates(list)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, getListVal(ret))
}

func Test_DuplicateList(t *testing.T) {
	var head *ListNode = new(ListNode)
	tail := head

	for i := 1; i <= 10; i++ {
		node := new(ListNode)
		if i%2 == 0 {
			node.Val = i - 1
		} else {
			node.Val = i
		}
		node.Next = nil

		tail.Next = node
		tail = node
	}

	ret := deleteDuplicates(head.Next)
	assert.Equal(t, []int{}, getListVal(ret))
}
