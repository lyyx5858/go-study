package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 { // 把所有利润为正的回起来就是最大利润
			max = max + prices[i] - prices[i-1]
		}
	}
	return max
}
