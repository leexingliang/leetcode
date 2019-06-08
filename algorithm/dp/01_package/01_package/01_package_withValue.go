package packageWithValue

// 我们刚刚讲的背包问题, 只涉及背包重量和物品重量。我们现在引入物品价值这一变量。
// 对于一组不同重量、不同价值、不可分割的物品，我们选择将某些物品装入背包，
// 在满足背包最大重量限制的前提下，背包中可装入物品的总价值最大是多少呢？

//  1. 回溯法
var maxV int                      // 结果
var number = 5                    // 物品数量
var maxW = 9                      // 背包最大承重
var weight = []int{2, 2, 4, 6, 3} // 物品重量
var values = []int{3, 4, 8, 9, 6} // 物品价值
func recursive(i int, cw int, cv int) {
	if i == number || cw == maxW {
		if cv > maxV {
			maxV = cv
			return
		}
	}
	// 不装
	recursive(i+1, cw, cv)
	// 装
	if cw+weight[i] <= maxW {
		recursive(i+1, cw+weight[i], cv+values[i])
	}
}

//

// 2. 记录状态 XXXXX不适合

//
var mem = [5][9 + 1]int{
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
}

func package01Iteration() int {

	mem[0][0], mem[0][weight[0]] = 0, values[0]

	for i := 1; i < number; i++ {
		// 不装
		for j := 0; j <= maxW; j++ {
			if mem[i-1][j] != -1 {
				mem[i][j] = mem[i-1][j]
			}
		}
		// 装
		for j := 0; j <= maxW-weight[i]; j++ {
			var val = mem[i-1][j] + values[i]
			if mem[i-1][j] != -1 || mem[i-1][j] < val {
				mem[i][j+weight[i]] = val
			}
		}
	}

	maxV = mem[number-1][0]
	for i := 1; i <= maxW; i-- {
		if mem[number-1][i] > maxV {
			maxV = mem[number-1][i]
		}
	}
	return maxV
}

//
