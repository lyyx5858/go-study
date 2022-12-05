package main

func main() {

	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 3, 4}
	list1 := buildList(arr1)
	list2 := buildList(arr2)
	list3 := mergeTwoLists(list1, list2)
	PrintList(list3)
}

var root *ListNode = &ListNode{}

func mergeTwoLists_liuyan(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	}
	cur := root
	for list1 != nil || list2 != nil {
		if list1.Val < list2.Val {
			cur.Val = list1.Val
			if list1.Next != nil {
				list1 = list1.Next
				cur.Next = &ListNode{}
				cur = cur.Next

			} else {
				cur.Next = list2
				break
			}
		} else {
			cur.Val = list2.Val
			if list2.Next != nil {
				list2 = list2.Next
				cur.Next = &ListNode{}
				cur = cur.Next

			} else {
				cur.Next = list1
				break
			}
		}
		PrintList(root)
	}

	return root
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
