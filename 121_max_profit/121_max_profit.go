package max_profit

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min := prices[0]
	max := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
			sub := max - min
			if profit < sub {
				profit = sub
			}
			continue
		}
		if prices[i] < min {
			min, max = prices[i], prices[i]
		}
	}
	return profit
}
