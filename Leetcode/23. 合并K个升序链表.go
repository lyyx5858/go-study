package main

func main() {

	list := make([]*ListNode, 3)
	a1 := []int{1, 4, 5}
	a2 := []int{1, 3, 4}
	a3 := []int{2, 6}
	list[0] = buildList(a1)
	list[1] = buildList(a2)
	list[2] = buildList(a3)
	l := mergeKLists(list)
	PrintList(l)

}

func mergeKLists(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	n := len(lists)
	if n < 1 {
		return nil
	}
	pre = lists[0]
	for i := 1; i < n; i++ {

		cur = lists[i]
		pre = mergeTwoLists(pre, cur)
	}
	return pre
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}

}
