package main

import "fmt"

func main() {
	arr := []int{3, 2, 0, -4}

	l1 := buildList(arr)
	l1.Next.Next.Next.Next = l1.Next //形成了环形
	fmt.Println(detectCycle(l1))

}

func detectCycle(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil {
		slow = slow.Next

		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow { //a:  直线，b:相交点, b+c=环长
			p := head
			for p != slow { //a=c+(n-1)(b+c), 当n=1时，a=c
				p = p.Next
				slow = slow.Next
			}
			return p
		}

	}

	return nil
}
