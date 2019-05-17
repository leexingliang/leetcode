package goods

import "fmt"

// 淘宝的“双十一”购物节有各种促销活动，比如“满 200 元减 50 元"
// 假设你女朋友的购物车中有 n 个（n>100）想买的商品，她希望从里面选几个，在凑够满减条件的前提下，
// 让选出来的商品价格总和最大程度地接近满减条件（200 元）

var goods = []int{20, 50, 33, 55, 98, 71}
var numbers = 6
var val = 200

// 物品价格, 物品数量, 满减条件
func choice(goods []int, numbers int, val int) { //[]int {
	fmt.Println("start...")
	state := new([6][3*200 + 1]bool)
	state[0][0], state[0][goods[0]] = true, true
	for i := 1; i < numbers; i++ {
		// 装
		for j := 0; j <= 3*val-goods[i]; j++ {
			if state[i-1][j] {
				state[i][j+goods[i]] = true
			}
		}
		// 不装
		for j := 0; j <= 3*val; j++ {
			if state[i-1][j] {
				state[i][j] = state[i-1][j]
			}
		}
	}
	// result := make([]int, 0)
	var index int
	for index = val; index <= 3*val; index++ {
		if state[numbers-1][index] {
			break
		}
	}
	fmt.Println(index)
	if index > 3*val {
		return
	}
	// for i := numbers - 2; i >= 0; i-- {
	// 	// 装了第 i+1 个物品
	// 	if state[i+1][index] && !state[i][index] {
	// 		fmt.Print(goods[i+1])
	// 		// result = append(result, goods[i+1])
	// 		index = index - goods[i+1]
	// 	}
	// }
	// if index > 0 {
	// 	fmt.Println(goods[0])
	// }

	for i := numbers - 1; i > 0; i-- {
		if index-goods[i] >= 0 && state[i-1][index-goods[i]] {
			fmt.Print(goods[i], " ")
			index = index - goods[i]
		}
	}
	if index > 0 {
		fmt.Println(goods[0])
	}

}
