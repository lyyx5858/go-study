package main

func main() {
	//a := []int{4, 2, 1, 3, 6}
	//a := []int{-1, 5, 3, 4, 0}
	//a := []int{}
	a := []int{4, 19, 14, 5, -3, 1, 8, 5, 11, 15}
	l := buildList(a)
	l2 := sortList(l)
	PrintList(l2)
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	v := middleLis(head) //每次取中间的结点的值
	l := &ListNode{}     //左
	m := &ListNode{}     //中间
	r := &ListNode{}     //右

	lc := l
	rc := r
	mc := m

	for head != nil {
		if head.Val < v {
			lc.Next = &ListNode{head.Val, nil}
			lc = lc.Next
		} else if head.Val > v {
			rc.Next = &ListNode{head.Val, nil}
			rc = rc.Next
		} else {
			mc.Next = &ListNode{head.Val, nil}
			mc = mc.Next
		}
		head = head.Next
	}

	l = sortList(l.Next)
	r = sortList(r.Next)

	return merge2List(merge2List(l, m.Next), r)

}

func merge2List(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	p := head1
	for p.Next != nil {
		p = p.Next
	}
	p.Next = head2
	return head1
}

func middleLis(head *ListNode) int {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow.Val
}
