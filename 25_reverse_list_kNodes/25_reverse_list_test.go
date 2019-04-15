package reverse_list_kNodes

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

func Test_ReverseList(t *testing.T) {
	list := createList(1, 5)

	ret := reverseKGroup(list, 2)

	assert.Equal(t, []int{2, 1, 4, 3, 5}, getListVal(ret))
}

func Test_ReverseList_EmptyList(t *testing.T) {
	list := createList(1, 0)

	ret := reverseKGroup(list, 2)

	assert.Equal(t, []int{}, getListVal(ret))
}

func Test_ReverseList_OneNodeList(t *testing.T) {
	list := createList(1, 1)

	ret := reverseKGroup(list, 2)

	assert.Equal(t, []int{1}, getListVal(ret))
}
