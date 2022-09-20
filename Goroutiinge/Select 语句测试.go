package main

import (
	"fmt"
	"time"
)

type li struct {
	ln       int
	connChan chan int
	errChan  chan error
}

func main() {

	li1 := li{
		ln:       0,
		connChan: make(chan int, 5),
		errChan:  make(chan error, 1),
	}

	//loop这个函数里有循环，
	//因此虽然这个函数只调用了一次
	//但只要main不结束，这个loop里需的循环
	//是一直在运行的。
	go li1.loop()

	for {
		r, err := li1.accept()
		fmt.Println("r is:", r, "err is:", err)
	}

	fmt.Println("main end....")
	time.Sleep(1 * time.Second)
}

func (li1 li) accept() (conn int, err error) {
	select {
	case conn = <-li1.connChan:
		fmt.Println(conn)
		fmt.Println("A")

	case err = <-li1.errChan:
		fmt.Println(err)
		fmt.Println("B")
	}
	return conn, err
}

func (li1 li) loop() {
	fmt.Println("loop start...")
	for i := 1; i < 10; i++ {
		//fmt.Println("I is:", i)
		li1.connChan <- i
		time.Sleep(1 * time.Second)
	}

}
