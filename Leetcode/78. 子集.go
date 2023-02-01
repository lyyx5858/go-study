package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

var (
	path []int
	res  [][]int
)

func subsets(nums []int) [][]int {
	path, res = make([]int, 0), make([][]int, 0)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, start int) {

	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for j := start; j < len(nums); j++ {

		path = append(path, nums[j])
		dfs(nums, j+1) //内循环，也叫深度循环, 前j个元素分析完了，开始分析剩下的
		path = path[:len(path)-1]
	}
}
