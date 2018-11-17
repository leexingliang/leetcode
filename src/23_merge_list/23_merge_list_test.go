package merge_list

import (
	"fmt"
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

func Test_MergeList(t *testing.T) {
	list0 := createList(1, 3)
	list1 := createList(4, 5)
	list2 := createList(6, 6)

	list := mergeKLists([]*ListNode{list0, list1, list2})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, getListVal(list))
}

func Test_MergeEmptyList(t *testing.T) {
	list0 := createList(1, 0)
	list1 := createList(1, 0)

	list := mergeKLists([]*ListNode{list0, list1})
	fmt.Println(list)
}

func Test_MergeList_fuck(t *testing.T) {
	list0 := new(ListNode)
	list0.Val = 2
	list0.Next = nil

	list1 := createList(1, 0)

	list2 := new(ListNode)
	list2.Val = -1
	list2.Next = nil

	list := mergeKLists([]*ListNode{list0, list1, list2})

	fmt.Println(list)
	assert.Equal(t, []int{-1, 2}, getListVal(list))

}

/// test leetcode solution

func Test_MergeList_leetcode(t *testing.T) {
	list0 := createList(1, 3)
	list1 := createList(4, 5)
	list2 := createList(6, 6)

	list := mergeKLists_leetcode([]*ListNode{list0, list1, list2})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, getListVal(list))
}

func Test_MergeEmptyList_leetcode(t *testing.T) {
	list0 := createList(1, 0)
	list1 := createList(1, 0)

	list := mergeKLists_leetcode([]*ListNode{list0, list1})
	fmt.Println(list)
}

func Test_MergeList_fuck_leetcode(t *testing.T) {
	list0 := new(ListNode)
	list0.Val = 2
	list0.Next = nil

	list1 := createList(1, 0)

	list2 := new(ListNode)
	list2.Val = -1
	list2.Next = nil

	list := mergeKLists_leetcode([]*ListNode{list0, list1, list2})

	fmt.Println(list)
	assert.Equal(t, []int{-1, 2}, getListVal(list))

}
