package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	//nums := []int{2, 0, 1, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	cover := 0
	n := len(nums) - 1
	for i := 0; i <= cover; i++ { // 每次与覆盖值比较
		cover = max(i+nums[i], cover) //每走一步都将 cover 更新为最大值
		if cover >= n {
			return true
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
