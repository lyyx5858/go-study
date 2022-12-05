package main

func main() {

	a := []int{3}
	l1 := buildList(a)
	b := []int{2}
	l2 := buildList(b)
	//PrintList(l1)
	//PrintList(l2)

	l2.Next = l1 //形成相交

	PrintList(l1)
	PrintList(l2)
	PrintList(getIntersectionNode(l1, l2))

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == headB {
		return headA
	}
	i := 0
	mapA := make(map[*ListNode]int)
	ha := headA
	hb := headB
	for ha != nil {
		mapA[ha] = i
		ha = ha.Next
		i++
	}

	for hb != nil {
		_, ok := mapA[hb]
		if ok {
			return hb
		}
		hb = hb.Next
	}

	return nil
}
