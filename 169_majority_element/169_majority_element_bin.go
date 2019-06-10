package majority_element

// 使用二分法好像不可以啊
// 无法确定边界条件

func majorityElement_bin(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for {
		index := binarry(nums, left, right)
		if index-left > right-index {
			right = index
		} else {
			left = index
		}

		if nums[left] == nums[right] {
			break
		}
	}
	return nums[left]
}

func binarry(nums []int, left, right int) int {
	max := compareAswap(nums, left, right)

	start, end := left, right
	for start < end {
		for start < end && nums[start] <= max {
			start++
		}
		nums[end] = nums[start]
		end--
		for start < end && nums[end] > max {
			end--
		}
		nums[start] = nums[end]
		start++
	}
	return start
}

func compareAswap(nums []int, left, right int) int {
	if nums[left] > nums[right] {
		nums[right], nums[left] = nums[left], nums[right]
	}
	return nums[right]
}
