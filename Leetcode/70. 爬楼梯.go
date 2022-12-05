package main

import "fmt"

func main() {
	n := 10
	fmt.Println(climbStairs(n))

}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i <= n-1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}
