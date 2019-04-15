package sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	lenght := len(nums) - 1
	for i := 0; i <= lenght-2; i++ {
		target := -nums[i]
		// 下标最大和最小值
		minI, maxI := i+1, lenght
		for minI < maxI {
			sum := nums[minI] + nums[maxI]
			switch {
			case sum > target:
				maxI--
			case sum < target:
				minI++
			case sum == target:
				result = append(result, []int{nums[i], nums[minI], nums[maxI]})
				// 去重
				for minI < maxI && nums[minI] == nums[minI+1] {
					minI++
				}
				for minI < maxI && nums[maxI] == nums[maxI-1] {
					maxI--
				}
				minI++
				maxI--
			}
		}
		// 去重
		for i < lenght && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}
