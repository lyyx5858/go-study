package main

import "fmt"

func main() {

	//v := []int{1}
	//v := []int{1, 2}
	v := []int{1, 2, 3, 4, 5}
	//v := []int{2, 3, -200, 1}
	//v := []int{2, 1, 4, 3, -200, 5}
	t := buildTree(v)

	fmt.Println(diameterOfBinaryTree(t))
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var dfs func(*TreeNode) int

	dfs = func(node *TreeNode) int { //dfs是个闭包函数
		if node == nil {
			return 0
		}

		// 计算每个结点左右孩子的高度
		left := dfs(node.Left)
		right := dfs(node.Right)

		// 路径长度等于左右孩子高度之和
		if left+right > res {
			res = left + right
		}

		// 返回较高的结点作为本结点的高度
		if left > right {
			return left + 1
		}

		return right + 1
	}

	dfs(root)
	return res
}
