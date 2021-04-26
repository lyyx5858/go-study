package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s *Student) SetName(name string) *Student {
	s.name = name
	return s
}

func (s *Student) SetAge(age int) *Student {
	s.age = age
	return s
}

func main() {
	s := Student{}
	s.SetName("li")
	s.SetAge(18)
	s.SetName("liu").SetAge(20)
	s.SetAge(21).SetName("Yang")
	fmt.Println(s)
}
