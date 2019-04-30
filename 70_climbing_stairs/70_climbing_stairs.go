package climbing_stairs

// func climbStairs(n int) int {
// 	if n <= 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}

// 	return climbStairs(n-1) + climbStairs(n-2)

// }

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	a, b := 1, 1
	for i := 2; i <= n; i++ {
		tmp := a
		a, b = b, tmp+b
	}
	return b
}
