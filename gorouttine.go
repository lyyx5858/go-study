package main

import "fmt"
import "time"

func hello() {
	fmt.Println("Hello Goroutine!")
}

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("Main Goroutine Done!")
}
