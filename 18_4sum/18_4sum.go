package sum4

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) <= 0 {
		return nil
	}

	sort.Ints(nums)

	end := len(nums) - 1
	result := make([][]int, 0)

	for i := 0; i <= end; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		array := numSum(nums, i, end, 4, target)
		for j := 0; j < len(array)-4; j += 4 {
			if array[j]+array[j+1]+array[j+2]+array[j+3] == 0 {
				result = append(result, array[j:j+4])
			}
		}
	}

	return result
}

func numSum(nums []int, start, end int, num int, sum int) [][]int {
	result := make([]int, 0)
	if start > end || num < end-start || num == 0 {
		return nil
	}
	// if num == 1 {
	// 	result = append(result, nums[start])
	// 	return result
	// }

	if num == 2 {

	}

	for i := start; i <= end; i++ {
		tmp := make([]int, 0, 4)
		tmp = append(tmp, nums[i])
		tmp = append(tmp, numSum(nums, start+1, end, num-1, sum-nums[start])...)
		if len(tmp) <= 4 {
			result = append(result, tmp...)
		}
	}

	return result
}

func sums(a []int) (sum int) {
	for _, val := range a {
		sum += val
	}
	return
}
