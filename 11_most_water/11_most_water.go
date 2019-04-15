package most_water

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	if len(height) == 1 {
		return height[0]
	}

	i, j := 0, len(height)-1
	water := 0

	for i < j {
		water = max(water, (j-i)*min(height[i], height[j]))
		// 如果 height[i] < height[j]
		// 说明 最大值有可能在 (i+1, j)之间
		// 否则 最大值有可能在 (i, j-1)之间
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return water
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
