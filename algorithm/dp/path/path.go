package path

// 假设我们有一个 n 乘以 n 的矩阵 w[n][n]。矩阵存储的都是正整数。
// 棋子起始位置在左上角，终止位置在右下角。我们将棋子从左上角移动到右下角。每次只能向右或者向下移动一位。
// 从左上角到右下角，会有很多不同的路径可以走。我们把每条路径经过的数字加起来看作路径的长度。
// 那从左上角移动到右下角的最短路径长度是多少呢？
// /////////////////或者也可以叫做杨辉三角问题

// 1. 回溯法
var min_dist int

// 第几行,第几列, 路径长度, 路径矩阵, 矩阵边界
func recursive(i int, j int, dist int, path [][]int, n int) {
	if i == n && j == n {
		if dist < min_dist {
			min_dist = dist
		}
	}
	// 往下走
	if i+1 < n {
		dist += path[i+1][j]
		recursive(i+1, j, dist, path, n)
	}
	// 往右走
	if j+1 < n {
		dist += path[i][j+1]
		recursive(i, j+1, dist, path, n)
	}
}

// 2. 动态规划方法
func recursiveMem(path [][]int, n int) int {

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = min(path[i-1][j]+path[i][j], path[i][j-1]+path[i][j])
		}
	}

	return path[n-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
