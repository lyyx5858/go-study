package main

func main() {

	a := []int{}
	l := buildList(a)
	PrintList(removeElements(l, 6))

}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}

func remove01(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return remove01(head.Next, val)
	}
	head.Next = remove01(head.Next, val)
	return head
}
