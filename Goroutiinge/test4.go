package main

import (
	"fmt"
	"time"
)

func main() {

	go loop_a()
	time.Sleep(time.Second * 15)

	fmt.Println("main end")

}

func loop_a() {
	fmt.Println("\tA start:")
	go loop_b()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("\t", i)
		if i == 9 {
			fmt.Println("\tA ends")
		}
	}

}

func loop_b() {
	fmt.Println("B start:")
	for i := 10; i < 20; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println(i)
	}

}
