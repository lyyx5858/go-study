package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func main() {
//
//	//v := []int{5, 4, 1, -200, 1, -200, 4, 2, -200, 2, -200}
//	//v := []int{3, 9, 20, -200, -200, 15, 7}
//	v := []int{1, 2, 2, 3, 4, 4, 3}
//	//v := []int{1, 2, 2, 2, -200, 2, -200}
//
//	t1 := buildTree(v)
//	fmt.Println()
//	fmt.Println(t1)
//	printTree(t1)
//
//}

func buildTree(data []int) (tree *TreeNode) {
	if len(data) == 0 {
		return nil
	}

	if isEven(len(data)) {
		data = append(data, -200)
	}

	que := []*TreeNode{}
	fill_left := true
	var node, root *TreeNode

	for _, d := range data {

		//先将数组data中的值构建成node
		if d != -200 {
			node = &TreeNode{Val: d}
		} else {
			node = nil
		}

		if len(que) == 0 {
			root = node             //先构建ROOT
			que = append(que, node) //然后将root node进队列

		} else if fill_left == true { //先加左
			que[0].Left = node //加在当前que中排在最前面的一个node的左边
			fill_left = false
			if node != nil { //非空的NODE进入队列，而只有放入队列的值才会加左、右 node
				que = append(que, node)
			}
		} else {
			que[0].Right = node //后加右
			if node != nil {
				que = append(que, node)
			}

			que = que[1:] //当右节点加上后，说明这个节点已经完成，因此从队列里取出
			fill_left = true
		}

	}

	return root
}

func calDepth(node *TreeNode) int {
	h := 0
	if node.Left != nil {
		h = calDepth(node.Left) + 1
	}
	if node.Right != nil {
		h = max(h, calDepth(node.Right)+1)
	}
	return h
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func isEven(num int) bool { //判断奇偶
	if num%2 == 0 {
		return true
	}
	return false
}

//以下是用来打印tree
func getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getTreeHeight(root.Left), getTreeHeight(root.Right))
}

func writeArray(root *TreeNode, row, column, treeHeight int, resArray [][]string) {
	if root == nil {
		return
	}
	resArray[row][column] = strconv.Itoa(root.Val)
	currentHeight := (row + 1) / 2 // 当前高度
	if currentHeight == treeHeight {
		return // 下面没有子节点了
	}
	gap := treeHeight - currentHeight - 1 // 到左/右儿子的距离
	// 填充左儿子
	if root.Left != nil {
		// 先写树结构符号
		resArray[row+1][column-gap] = "/"
		// 再写左儿子
		writeArray(root.Left, row+2, column-gap*2, treeHeight, resArray)
	}
	// 填充右儿子
	if root.Right != nil {
		resArray[row+1][column+gap] = "\\"
		writeArray(root.Right, row+2, column+gap*2, treeHeight, resArray)
	}
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("The tree is nil!")
		return
	}
	height := getTreeHeight(root)
	fmt.Printf("height: %v\n", height)
	// 总宽度为节点高度 * 2 - 1, 因为还要画树枝符号
	totalHeight := height*2 - 1
	// 最大宽度为3 * 2^(n - 1) + 1，公式如下：
	/**
	   父亲节点占1, 两个孩子空间各占1, 连接线各占1, 每个父子关系共占5, 每个关系之间空1, 最左最右各空1
	  第2行： 5 + 2 （1个父子结构占位+左右两个空格分割）
	  第3行：2 * 5 + (1 + 2) （2个父子结构占位+1个相邻父子结构间空格+左右两个空格分割）
	  第4行：4 * 5 + (3 + 2) （4个父子结构占位+3个相邻父子结构间空格+左右两个空格分割）
	  第5行：8 * 5 + (7 + 2)
	  第n行：5 * 2 ^ (n - 2) + (2 ^ (n - 2) - 1) + 2 = 6 * 2 ^ (n-2) + 1 = 3 * 2 ^ (n - 1) + 1
	*/
	var totalWidth int
	if height == 0 {
		totalWidth = 1
	} else {
		totalWidth = (2<<(height-2))*3 + 1
	}

	// 创建数组
	printArray := make([][]string, totalHeight)
	for i := range printArray {
		printArray[i] = make([]string, totalWidth)
		for j := range printArray[i] {
			printArray[i][j] = " "
		}
	}

	// 计算打印数组
	writeArray(root, 0, totalWidth/2, height, printArray)

	// 打印
	for i := range printArray {
		var res string
		for j := range printArray[i] {
			res = res + printArray[i][j]
		}
		fmt.Println(res)
	}
}
