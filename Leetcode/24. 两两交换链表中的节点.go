package main

func main() {
	a := []int{1, 2, 3, 4}
	l := buildList(a)

	PrintList(swapPairs01(l))

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next
		cur.Next = cur.Next.Next
		tmp.Next = tmp.Next.Next
		cur.Next.Next = tmp
		cur = cur.Next.Next
	}

	return dummy.Next

}

func swapPairs01(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newhead := head.Next
	head.Next = swapPairs01(newhead.Next)
	newhead.Next = head
	return newhead
}
