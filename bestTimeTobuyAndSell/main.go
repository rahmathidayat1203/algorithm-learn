package main

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit(prices)
}

func maxProfit(prices []int) int {
	profit := 0
	i := 1
	for i < len(prices) {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
		i++
	}
	return profit
}
