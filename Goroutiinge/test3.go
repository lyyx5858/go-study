package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1
	go func() {
		v := <-ch
		go fmt.Println(v)
	}()
	fmt.Println("2")
	time.Sleep(1 * time.Second)

}
