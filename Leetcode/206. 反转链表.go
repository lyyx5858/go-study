package main

func main() {
	a := []int{1, 2, 3, 4, 5}
	l := buildList(a)

	PrintList(l)
	PrintList(reverseList00(l))

}

var res *ListNode
var cur = &ListNode{}
var h0 int

func reverseList00(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res = cur

	if head.Next != nil {
		h0++
		reverseList00(head.Next)
	}

	h0--
	cur.Val = head.Val
	if h0 >= 0 {
		cur.Next = &ListNode{}
	}
	cur = cur.Next

	return res
}
func reverseList01(head *ListNode) *ListNode {
	var res0 *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = res0 //将现有的head的下一跳置空
		res0 = head
		head = tmp
	}

	return res0

}
func reverseList02(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	new := reverseList02(head.Next) //假如head目前是2， 则这句会让3 4 5 反转完成。
	head.Next.Next = head           //1->2->三<-4<-5 ,2的Next是3，这句话的意思是让3的Next为2
	head.Next = nil                 //然后让2的下跳为空，不然会循环
	return new

}
