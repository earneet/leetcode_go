package Problem_101_200

func maxProfit(prices []int) int {
	max := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}

		diff := prices[i] - min
		if diff > max {
			max = diff
		}
	}
	return max
}
