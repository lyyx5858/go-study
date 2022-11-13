package main

import "fmt"

func main() {
	root1 := []int{1, 3, 2, 5}
	root2 := []int{2, 1, 3, -200, 4, -200, 7}

	r1 := buildTree(root1)
	r2 := buildTree(root2)
	r3 := mergeTrees(r1, r2)
	fmt.Println(r3)

}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	root := &TreeNode{}
	if root1 == nil {
		root = root2
	} else if root2 == nil {
		root = root1
	} else {
		root.Val = root1.Val + root2.Val
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
	}

	return root
}
