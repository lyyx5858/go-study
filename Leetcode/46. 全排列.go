package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3}
	fmt.Println(permute(nums1))

	nums2 := []int{0, 1}
	fmt.Println(permute(nums2))
}

var (
	path []int
	res  [][]int
)

func permute(nums []int) [][]int {
	path, res = make([]int, 0), make([][]int, 0)
	dfs(nums, len(nums)) //将原始长度传进去，以用来判断记录条件
	return res
}

func dfs(nums []int, l int) {

	if len(path) == l { //当path等于nums长度时，表示排列完成
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}

	for j := 0; j < len(nums); j++ {
		t := []int{} //临时生成一个数组
		path = append(path, nums[j])
		t = append(append(t, nums[:j]...), nums[j+1:]...) //将nums[j]排除后，继续循环，排列是可以回头的，但不能重复
		dfs(t, l)                                         //内循环，也叫深度循环, 第j个元素分析完了，开始分析剩下的
		path = path[:len(path)-1]
	}
}
