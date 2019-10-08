package search_insert_position

func searchInsert(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}

	var i int
	for i = 0; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
	}
	return i
}
