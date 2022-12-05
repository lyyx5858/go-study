package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 0, -4}
	l1 := buildList(arr)
	cru := l1
	cru.Next.Next.Next.Next = cru.Next //形成了环形
	fmt.Println(hasCycle(l1))

}

var m = make(map[*ListNode]int)
var i int = 1

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	_, ok := m[head] // 看表中是否已有节点

	if ok == false {
		m[head] = i //如果表中没有节点，则创建一项
		i++
		return hasCycle(head.Next)
	} else {
		return true
	}
}
