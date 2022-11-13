package main

func main() {

	arr1 := []int{1, 5, 6}
	arr2 := []int{2, 3, 4}
	list1 := buildList(arr1)
	list2 := buildList(arr2)
	list3 := mergeTwoLists(list1, list2)
	PrintList(list3)
}

var new *ListNode = &ListNode{}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		new = addEndnode(new, list1)
		list1 = list1.Next
		new = addEndnode(new, mergeTwoLists(list1.Next, list2))
	} else {
		new = addEndnode(new, list2)
		list2 = list2.Next
		new = addEndnode(new, mergeTwoLists(list1, list2.Next))
	}

	return new
}

func addEndnode(root *ListNode, list *ListNode) *ListNode {
	root.Next = list
	return root
}
