package main

import (
    "fmt"
)

type Employee struct {
    name string
    age  int
}

//type Unicom struct {
//    Em  []Employee
//
//}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
    e.name = newName
	fmt.Printf(e.name)
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
    e.age = newAge
}

func main() {
    e := Employee{
        name: "Mark Andrew",
        age:  50,
    }

    var Em []Employee

  //  Em=make([]Employee,2)
    Em[0]=e
    // Em[1] :=e

    fmt.Printf("\nEmployee name before change: %s\n", e.name)
    e.changeName("liu")
    fmt.Printf("\nEmployee name after change: %s", e.name)

    fmt.Printf("\n\nEmployee age before change: %d", e.age)
    e.changeAge(51)
    fmt.Printf("\nEmployee age after change: %d\n", Em[0].age)
   // fmt.Printf("\nEmployee age after change: %d\n", Em)


}
