package merge_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 每个链表去一个元素, 构建一个最小堆
// 每次取堆顶的元素插入新的链表,
// 并且将被取出元素的下一个元素插入最小堆
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	var head *ListNode = new(ListNode)
	var tail *ListNode = head
	var minHeap []*ListNode = make([]*ListNode, 0, length)
	for _, linklist := range lists {
		if linklist != nil {
			minHeap = append(minHeap, linklist)
		}
	}

	// 构建最小堆
	heapLength := len(minHeap)
	for len(minHeap) > 1 {
		createMinHeap(minHeap, heapLength)
		for minHeap[0].Next != nil {
			tail.Next = minHeap[0]
			tail = tail.Next
			minHeap[0] = minHeap[0].Next
			adjustHeapSort(minHeap, 0, heapLength)
		}
		// 删掉第一个节点
		tail.Next = minHeap[0]
		tail = tail.Next
		minHeap = minHeap[1:]
		heapLength--
	}
	if len(minHeap) > 0 {
		tail.Next = minHeap[0]
	}

	return head.Next
}

func createMinHeap(array []*ListNode, len int) {
	for i := len/2 - 1; i >= 0; i-- {
		// 从第一个非叶子节点从下至上, 从右向左调整
		adjustHeapSort(array, i, len)
	}
}

func adjustHeapSort(array []*ListNode, index, len int) {
	node := array[index]
	for i := 2*index + 1; i < len; i = 2*i + 1 {
		// 如果右叶子节点值小于左叶子节点值
		if i+1 < len && array[i+1].Val < array[i].Val {
			i++
		}
		if array[i].Val < node.Val {
			array[index] = array[i]
			index = i
		} else {
			// 不需要调整
			break
		}
	}
	array[index] = node
}
