package main

import "fmt"

func main() {

	root := []int{1, 2, 5, 3, 4, -200, 6}
	//root := []int{1, 2, 5}
	r1 := buildTree(root)
	flatten(r1)
	fmt.Println(r1)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left) //先将左边展开

	if root.Left != nil {
		l := root.Left       //l为临时用
		for l.Right != nil { //找到l的最下边节点
			l = l.Right
		}
		l.Right = root.Right   //将root右边接到l的最下边节点
		root.Right = root.Left //将整个右节点搬过来
		root.Left = nil        //将左边置空
	}
	flatten(root.Right) //同样处理右边节点

}
