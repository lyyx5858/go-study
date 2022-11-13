package main

import "fmt"

func main() {

	//v := []int{1}
	v := []int{1, 2}
	//v := []int{1, 2, 3, 4, 5}
	//v := []int{2, 3, -200, 1}
	//v := []int{2, 1, 4, 3, -200, 5}
	//v := []int{3, 9, 20, -200, -200, 15, 7}
	t := buildTree(v)
	fmt.Println(maxDepth(t))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := dfs(node.Left) + 1
		r := dfs(node.Right) + 1

		return max(l, r)
	}
	res = dfs(root)
	return res
}
