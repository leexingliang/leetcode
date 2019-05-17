package package_01

// 我们有一个背包，背包总的承载重量是 Wkg。
// 现在我们有 n 个物品，每个物品的重量不等，并且不可分割。
// 我们现在期望选择几件物品，装载到背包中。
// 在不超过背包所能装载重量的前提下，如何让背包中物品的总重量最大？

// 1. 回溯算法实现
var maxValue int                  // 最大重量结果
var number = 5                    // 物品个数
var maxWeight = 9                 // 背包最大承受重量
var weight = []int{2, 2, 4, 6, 3} // 物品重量
func package01Recursive(i int, cw int) { // 参数是当前物品下标, 和当前物品重量
	// 判断是否所有物品都已经装过, 或者装入的物品超过了背包最大重量
	if i == number || cw == maxWeight {
		// 保存结果
		if cw > maxValue {
			maxValue = cw
		}
	}

	// 不装这个物品
	package01Recursive(i+1, cw)
	// 装这个物品
	if cw+weight[i] <= maxWeight {
		package01Recursive(i+1, cw+weight[i])
	}
}

// 2. 进一步优化, 在每次的计算中都会有重复的计算, 如下图
// 									f(0, 0)
// 					f(1, 0) 									f(1, 2)
// 		 f(2, 0) 			f(2, 2) 				f(2, 2) 				f(2, 4)
//  f(3, 0)   f(3, 4)    f(3, 2)  f(3, 6)        f(3, 0)  f(3, 6)       f(3, 4)  f(3, 8)
//    ..............
// 我们可以借助一个数组, 存储中间转态
var mem = [5][9]bool{}

func package01Mem(i int, cw int) {
	if i == number || cw == maxWeight {
		if cw > maxValue {
			maxValue = cw
		}
	}
	if mem[i][cw] { // 重复状态
		return
	}

	mem[i][cw] = true // 记录状态 ********
	// 不装这个物品
	package01Recursive(i+1, cw)
	// 装这个物品
	if cw+weight[i] <= maxWeight {
		package01Recursive(i+1, cw+weight[i])
	}
}

// 3. 我们可以在简化一点, 使用迭代消除递归带来的复杂度
func package01Iteration() int {

	var mem = new([5][9 + 1]bool)
	// 初始化状态
	mem[0][0], mem[0][weight[0]] = true, true

	for i := 1; i < number; i++ {
		// 装
		for j := 0; j < maxWeight; j++ {
			if mem[i-1][j] && j+weight[i] <= maxWeight {
				mem[i][j+weight[i]] = true
			}
		}

		// 不装
		for j := 0; j < maxWeight; j++ {
			if mem[i-1][j] {
				mem[i][j] = mem[i-1][j]
			}
		}
	}

	for i := maxWeight; i > 0; i-- {
		if mem[number][i] == true {
			return i
		}
	}
	return 0
}

// 4. 简化状态数组, 将二维数组替换为一维数组
func package01Iteration_Final() int {

	var mem = new([9 + 1]bool)
	// 初始化状态
	mem[0], mem[weight[0]] = true, true

	for i := 1; i < number; i++ {
		// 装

		for j := maxWeight - weight[i]; j >= 0; j-- {
			if mem[j] {
				mem[j+weight[i]] = true
			}
		}
		// 从小到大计算会出现重复问题
		// for j := 0; j < maxWeight; j++ {
		// 	if mem[j] && j+weight[i] <= maxWeight {
		// 		mem[j+weight[i]] = true
		// 	}
		// }
		//----------------------------------------------//

		// 不装, 就不用管了, 复用上一次的就行了
		// for j := 0; j < maxWeight; j++ {
		// 	if mem[i-1][j] {
		// 		mem[i][j] = mem[i-1][j]
		// 	}
		// }
	}

	for i := maxWeight; i > 0; i-- {
		if mem[i] == true {
			return i
		}
	}
	return 0
}
