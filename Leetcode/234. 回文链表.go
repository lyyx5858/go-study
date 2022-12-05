package main

import "fmt"

func main() {
	a := []int{1, 1}
	l1 := buildList(a)

	fmt.Println(isPalindrome(l1))
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	var r bool
	if head.Val == lastNode(head).Val {
		r = isPalindrome(cutlastNode(head.Next))
	} else {
		r = false
	}

	return r
}

func lastNode(head *ListNode) *ListNode {

	for head.Next != nil {
		head = head.Next
	}
	return head
}

func cutlastNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	return head
}
