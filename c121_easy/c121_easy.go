package c121_easy

import "math"

func maxProfit(prices []int) int {
	minprice, maxprofit := math.MaxInt32, 0
	for _, val := range prices {
		if val < minprice {
			minprice = val
		} else if val - minprice > maxprofit {
			maxprofit = val - minprice
		}
	}
	return maxprofit
}
