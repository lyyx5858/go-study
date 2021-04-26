package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("1")

	fmt.Println("2")
	time.Sleep(1 * time.Second)

}
