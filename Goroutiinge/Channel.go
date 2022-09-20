package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	//ch <- 1

	//loop这个函数里有循环，
	//因此虽然这个函数只调用了一次
	//但只要main不结束，这个loop里需的循环
	//是一直在运行的。
	go loop(ch)

	//for {
	//	go accept(ch)
	//}

	<-ch
	<-ch
	<-ch

	fmt.Println("main end....")
	time.Sleep(1 * time.Second)
}

func accept(ch chan int) {
	select {
	case v := <-ch:
		fmt.Println(v)
		fmt.Println("A")
	}
}

func loop(ch chan int) {
	fmt.Println("loop start...")

	for i := 1; i < 10; i++ {
		fmt.Println("I is:", i)
		ch <- i
		time.Sleep(1 * time.Second)
	}

}
