package main

import (
	"fmt"
)

func main() {
	root := []int{3, 5, 1, 6, 2, 0, 8, -200, -200, 7, 4}
	p := []int{5, 6, 2, -200, -200, 7, 4}
	q := []int{1, 0, 8}
	r1 := buildTree(root)
	p1 := buildTree(p)
	q1 := buildTree(q)
	r2 := lowestCommonAncestor(r1, p1, q1)
	fmt.Println(r2)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)   //在root的左边找p或者q，只要有任意一个，left的值就不会是nil
	right := lowestCommonAncestor(root.Right, p, q) //在root的右边找p或者q，只要有任意一个，right的值就不会是nil
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func findNode(root *TreeNode, x int) (res *TreeNode) {

	if root == nil {
		return nil
	} else if root.Val == x {
		fmt.Println(root.Val)
		return root
	} else {
		fmt.Println(root.Val)
	}

	if root.Left != nil {
		if root.Left.Val == x {
			return root.Left
		} else {
			res = findNode(root.Left, x)
		}
	}

	if res == nil {
		if root.Right != nil {
			if root.Right.Val == x {
				return root.Right
			} else {
				res = findNode(root.Right, x)
			}
		}
	}
	return res
}
