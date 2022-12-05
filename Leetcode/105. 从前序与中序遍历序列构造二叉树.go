package main

func main() {
	//v := []int{3, 9, 20, -200, -200, 15, 7}
	v := []int{3, 9, 20, 1, 2, 15, 7}
	t := buildTree(v)
	printTree(t)

	preorder := []int{3, 9, 1, 2, 20, 15, 7}
	inorder := []int{1, 9, 2, 3, 15, 20, 7}
	t1 := buildTree_from_pre_in(preorder, inorder)
	printTree(t1)

}

//Preorder: 根左右，根在前，先序遍历
//inorder: 	左根右，根在中间，中序遍历
//postorder: 左右根，根在后面，后序遍历

func buildTree_from_pre_in(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree_from_pre_in(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree_from_pre_in(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}
