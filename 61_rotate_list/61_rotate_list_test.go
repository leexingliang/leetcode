package rotate_list

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

func Test_RotateList_One(t *testing.T) {
	list := createList(1, 5)

	ret := rotateRight(list, 1)
	assert.Equal(t, []int{5, 1, 2, 3, 4}, getListVal(ret))
}

func Test_RotateList_Two(t *testing.T) {
	list := createList(1, 5)

	ret := rotateRight(list, 2)
	assert.Equal(t, []int{4, 5, 1, 2, 3}, getListVal(ret))
}
func Test_RotateList_Three(t *testing.T) {
	list := createList(1, 5)

	ret := rotateRight(list, 3)
	assert.Equal(t, []int{3, 4, 5, 1, 2}, getListVal(ret))
}

func Test_RotateList_ListLen(t *testing.T) {
	list := createList(1, 5)

	ret := rotateRight(list, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, getListVal(ret))
}

func Test_RotateList_ListLenAdd1(t *testing.T) {
	list := createList(1, 5)

	ret := rotateRight(list, 6)
	assert.Equal(t, []int{5, 1, 2, 3, 4}, getListVal(ret))
}
