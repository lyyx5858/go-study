package main

import (
	"fmt"
	"sort"
)

func main() {
	canidates := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	canidates1 := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(canidates, 1))
	fmt.Println(combinationSum2(canidates1, 8))
}

var (
	sum  int
	res  [][]int
	path []int
)

func combinationSum2(candidates []int, target int) [][]int {
	path = make([]int, 0)
	res = make([][]int, 0)
	dfs(candidates, target, 0)
	return res
}

func dfs(candidates []int, target int, sum int) {
	sort.Ints(candidates) //排序
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	} else if sum > target {
		return
	}

	for j := 0; j < len(candidates); j++ { //外循环

		if j > 0 { //第0个元素不进入此判断
			for candidates[j] == candidates[j-1] && j < len(candidates)-1 { //与前一元素相同，且不是最后一个
				j++
			}
			if candidates[j] == candidates[j-1] && j == len(candidates)-1 { //与前一元素相同，是最后一个
				break
			}
		}

		path = append(path, candidates[j])
		dfs(candidates[j+1:], target, sum+candidates[j]) //内循环，从j+1开始，不回头
		path = path[:len(path)-1]
	}
}
