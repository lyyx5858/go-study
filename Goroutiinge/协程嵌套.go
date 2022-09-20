package main

import (
	"fmt"
	"time"
)

func main() {
	go func() { //父协程
		fmt.Println("______父协程 Start______")

		go func() { //子协程
			for i := 0; i < 10; i++ {
				time.Sleep(time.Second * 2)
				fmt.Println("   子协程", i)
			}
		}()

		time.Sleep(time.Second * 4)
		fmt.Println("______父协程 End______") //父协程已经结束的情况下，子协程依然进行
		//但如果main结束，任何协程都结束
		return
	}()

	time.Sleep(time.Second * 10)

	fmt.Println("\n======Main is 结束 ========")
}
