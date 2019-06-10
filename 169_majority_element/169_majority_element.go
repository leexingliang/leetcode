package majority_element

func majorityElement(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	var num = 1
	var data = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == data {
			num++
		} else {
			if num > 0 {
				num--
			} else {
				data = nums[i]
				num = 1
			}
		}
	}
	return data
}
