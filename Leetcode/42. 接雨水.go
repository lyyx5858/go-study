package main

import "fmt"

func main() {
	//height := []int{0, 1, 0, 2, 1, 0, 1, 30, 2, 1, 2, 1}
	height := []int{3}
	fmt.Println(trap(height))

}

func trap(height []int) int {
	ans := 0
	l := len(height)
	n := 0
	hmax := height[0]
	for i := 0; i <= l-1; i++ { //先找出最高的墙
		if height[i] > hmax {
			hmax = height[i]
			n = i
		}
	}

	dp := make([]int, l) //每个墙能接雨水的高度
	dp[0] = height[0]
	dp[l-1] = height[l-1]
	for i := 1; i < n; i++ { //计算最高墙左边可接的雨水总量，当最高墙在右边时，能接雨水的量只取决于左边墙高度

		if height[i] > dp[i-1] {
			dp[i] = height[i]
		} else {
			dp[i] = dp[i-1]
		}
		ans = ans + (dp[i] - height[i]) //能接雨水高度-墙高=实际能接雨水的量
	}

	for i := l - 2; i > n; i-- { //计算最高墙右边可接的雨水总量， 当最高墙在左边时，能接雨水的量只取决于右边墙高度

		if height[i] > dp[i+1] {
			dp[i] = height[i]
		} else {
			dp[i] = dp[i+1]
		}
		ans = ans + (dp[i] - height[i])

	}

	return ans
}
