package main

import "fmt"

func main() {

	queue := Constructor()
	fmt.Println(queue)
	queue.Push(1)
	fmt.Println(queue)
	queue.Push(2)
	fmt.Println(queue)
	fmt.Println(queue.Peek()) // 返回 1
	fmt.Println(queue)
	fmt.Println(queue.Pop()) // 返回 1
	fmt.Println(queue)
	queue.Empty() // 返回 false
	fmt.Println(queue)

}

type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}
