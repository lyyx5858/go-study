package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//func main() {
//
//	arr := []int{1, 2, 4}
//	l := buildList(arr)
//	PrintList(l)
//
//}

func buildList(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	root := &ListNode{Val: arr[0]}
	cur := root
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}

	return root
}

func PrintList(root *ListNode) {
	if root == nil {
		fmt.Printf("\n")
		return
	}
	if root.Next != nil {
		fmt.Printf("%d-->", root.Val)
	} else {
		fmt.Printf("%d", root.Val)
	}

	PrintList(root.Next)
}
