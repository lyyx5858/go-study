package main

import (
	"fmt"
)

func main() {
	canidates := []int{2}
	fmt.Println(combinationSum(canidates, 1))
}

var (
	sum  int
	res  [][]int
	path []int
)

func combinationSum(candidates []int, target int) [][]int {
	path = make([]int, 0)
	res = make([][]int, 0)
	dfs(candidates, target, 0)
	return res
}

func dfs(candidates []int, target int, sum int) {
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	} else if sum > target {
		return
	}

	for j := 0; j < len(candidates); j++ { //外循环

		path = append(path, candidates[j])
		dfs(candidates[j:], target, sum+candidates[j]) //内循环，从j开始，不回头
		path = path[:len(path)-1]
	}
}
