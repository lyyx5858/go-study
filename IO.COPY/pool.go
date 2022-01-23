package main

import (
	"fmt"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return "123"
	},
}

func main() {
	t := pool.Get().(string)
	fmt.Println(t)

	pool.Put("321")

	t2 := pool.Get().(string)
	fmt.Println(t2)
}
