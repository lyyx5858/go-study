package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

var (
	path []int
	res  [][]int
)

func subsetsWithDup(nums []int) [][]int {
	path, res = make([]int, 0), make([][]int, 0)
	sort.Ints(nums)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for j := start; j < len(nums); j++ {
		if j != start && nums[j] == nums[j-1] { //只有横向循环才有可能有重复。因此，当j比start大，且有重复无此时，剪枝。
			continue //别忘了nums是已经排序过了的。
		}

		path = append(path, nums[j])
		dfs(nums, j+1) //内循环，也叫深度循环, 前j个元素分析完了，开始分析剩下的
		path = path[:len(path)-1]

	}
}
