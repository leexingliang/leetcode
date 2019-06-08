package house_robber

func rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	prev := 0
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := curr
		curr = max(curr, prev+nums[i])
		prev = tmp
	}
	return curr
}

func rob_2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	mem := make([]int, len(nums)+1)

	mem[0] = 0
	mem[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		mem[i+1] = max(mem[i], mem[i-1]+nums[i])
	}
	return mem[len(nums)]
}

// var mem []int

// func rob_1(nums []int) int {
// 	if len(nums) <= 0 {
// 		return 0
// 	}
// 	mem = make([]int, len(nums)+1)
// 	for i := 0; i < len(mem); i++ {
// 		mem[i] = -1
// 	}
// 	return tryRob(nums, len(nums)-1)
// }

// func tryRob(nums []int, pos int) int {
// 	if pos < 0 {
// 		return 0
// 	}
// 	if mem[pos] != -1 {
// 		return mem[pos]
// 	}
// 	tmp := max(tryRob(nums, pos-2)+nums[pos], tryRob(nums, pos-1))
// 	mem[pos] = tmp
// 	return tmp
// }

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
