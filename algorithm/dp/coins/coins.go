package coins

import (
	"math"
)

// 硬币找零问题
// 我们今天来看一个新的硬币找零问题。假设我们有几种不同币值的硬币 v1，v2，……，vn（单位是元）。
// 如果我们要支付 w 元，求最少需要多少个硬币。
// 比如，我们有 3 种不同的硬币，1 元、3 元、5 元，我们要支付 9 元，最少需要 3 个硬币（3 个 3 元的硬币）。
var coins = []int{1, 3, 5} // 硬币集合
var pay = 9                // 支付多少元
var num = math.MaxInt32    //结果

//  1. 回溯法
func calcCoins(coins []int, currPay int, currNum int) {
	if currPay > pay {
		return
	}
	if currPay == pay {
		if currNum < num {
			num = currNum
		}
		return
	}

	// 装一个 1 元的
	calcCoins(coins, currPay+1, currNum+1)
	// 装一个 3 元的
	calcCoins(coins, currPay+3, currNum+1)
	// 装一个 5 元的
	calcCoins(coins, currPay+5, currNum+1)
}

// 2. 回溯记忆法
// 行: 最小数额的硬币组成支付币值得数量
// 列: 币值
var mem [9 + 1][9 + 1]bool

func calcCoinsMem(coins []int, currPay int, currNum int) {
	if currPay == pay {
		if currNum < num {
			num = currNum
		}
		return
	}

	mem[currNum][currPay] = true
	for i := 0; i < len(coins); i++ {
		if currPay+coins[i] <= pay && !mem[currNum+1][currPay+coins[i]] {
			calcCoinsMem(coins, currPay+coins[i], currNum+1)
		}
	}
	// // 装一个 1 元的
	// if currPay+coins[i] <= pay && !mem[currNum+1][currPay+1] {
	// 	calcCoinsMem(coins, currPay+1, currNum+1)
	// }
	// // 装一个 3 元的
	// if currPay+coins[i] <= pay && !mem[currNum+1][currPay+3] {
	// 	calcCoinsMem(coins, currPay+3, currNum+1)
	// }
	// // 装一个 5 元的
	// if currPay+coins[i] <= pay && !mem[currNum+1][currPay+5] {
	// 	calcCoinsMem(coins, currPay+5, currNum+1)
	// }
}

// 3. 迭代的方式
func calcCoinsF(coins []int) int {
	var state [9 + 1][9 + 1]bool
	for i := 0; i < len(coins); i++ {
		state[1][coins[i]] = true
	}
	for i := 1; i < 9; i++ {
		for j := 0; j <= pay; j++ {
			if state[i][j] {
				for k := 0; k < len(coins); k++ {
					if j+coins[k] <= pay {
						state[i+1][j+coins[k]] = true
					}
				}
			}
		}
	}
	for i := 1; i <= 9; i++ {
		if state[i][pay] {
			return i
		}
	}
	return 0
}
