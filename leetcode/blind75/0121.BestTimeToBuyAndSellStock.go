package blind75

func maxProfit(prices []int) int {
	max := 0
	min := 1<<31 - 1
	for _, price := range prices {
		if price < min {
			min = price
		}
		if max < price-min {
			max = price - min
		}
	}
	return max
}
