package main

import (
	"fmt"
	"math/rand"
)

type LuckyNumber struct {
	number int
}

type Person struct {
	lucky_numbers []LuckyNumber
}

func main() {

	tmp := make([]LuckyNumber, 10)
	for i := range tmp {
		tmp[i].number = rand.Intn(100)
	}
	a := Person{tmp}
	fmt.Println(a)

}