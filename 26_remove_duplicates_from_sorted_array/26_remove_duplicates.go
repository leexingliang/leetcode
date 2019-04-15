package remove

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	length := 1
	var value = nums[0] //
	var index = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != value {
			value = nums[i]
			nums[index] = value
			index++
			length++
		}
	}
	return length
}
