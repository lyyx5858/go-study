package main

import "fmt"

func main() {

	//nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	nums := []int{2, 0, 1, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	m := len(nums)
	dp := make([]bool, m)
	dp[m-1] = true
	for i := m - 2; i >= 0; i-- {
		n := min(nums[i], m-i)
		for j := 1; j <= n; j++ {
			if dp[i+j] == true {
				dp[i] = dp[i+j]
				break
			} else {
				dp[i] = false
			}
		}
	}

	fmt.Println(dp)
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
