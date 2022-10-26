package main

import "fmt"

func main() {
	nums := []int{-1, -2}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return nums[0]
	}
	dp := make([]int, l) //dp的定义：截至到目前子串的加和的最大值
	max := nums[0]
	dp[0] = nums[0]
	for i := 1; i <= l-1; i++ {
		dp[i] = dp[i-1] + nums[i]
		if dp[i] < nums[i] { //如果加上当前nums[i]反而小(上行执行结果），新更新dp[i]为nums[i]
			dp[i] = nums[i]
		}

		if dp[i] > max {
			max = dp[i]
		}

	}

	return max
}
