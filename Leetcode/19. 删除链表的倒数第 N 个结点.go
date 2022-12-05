package main

func main() {
	a := []int{1, 2}
	l := buildList(a)

	l1 := removeNthFromEnd(l, 2)
	PrintList(l1)

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	head = &ListNode{0, head}
	n = n + 1

	p := head
	q := head

	for i := 0; i < n; i++ {
		p = p.Next
	}

	for p != nil {
		p = p.Next
		q = q.Next
	}

	q.Next = q.Next.Next

	return head.Next
}
