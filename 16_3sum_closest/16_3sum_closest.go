package sum

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var result = nums[0] + nums[1] + nums[2]
	lenght := len(nums) - 1
	for i := 0; i <= lenght-2; i++ {
		// 下标最大和最小值
		minI, maxI := i+1, lenght
		for minI < maxI {
			sum := nums[minI] + nums[maxI] + nums[i]
			if sub(sum, target) < sub(result, target) {
				result = sum
			}
			if sum < target {
				minI++
			} else if sum > target {
				maxI--
			} else {
				return target
			}
		}
	}
	return result
}

func sub(a, b int) int {
	value := a - b
	if value < 0 {
		value = -value
	}
	return value
}
