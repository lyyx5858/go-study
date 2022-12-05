package main

import "fmt"

func main() {
	p := []int{1, 7}
	//p := []int{1, 6, 8, 1, 0, 12, 3, 1, 13, 11}
	fmt.Println(maxProfit(p))
}

func maxProfit(prices []int) int {
	l := len(prices)
	if l == 1 {
		return 0
	}
	dp := make([]int, l) //此处dp[i]代表截至前i个股票值，计算出来的最大利润
	pmin := prices[0]    //始终记录最便宜的价格
	dp[0] = 0            //初始值为0
	for i := 1; i <= l-1; i++ {
		if prices[i] < pmin {
			pmin = prices[i] //不断更新最便宜的价格
		}

		//dp[i]=max(prices[i]-pmin, dp[i-1])
		if prices[i]-pmin > dp[i-1] { //如果出现了最大利润，则更利润矩阵
			dp[i] = prices[i] - pmin
		} else {
			dp[i] = dp[i-1] //反之当前利润值等于以前
		}
	}

	return dp[l-1]
}
