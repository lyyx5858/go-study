package main

import (
    "fmt"
    "flag"
)

func main() {
    b := flag.Bool("b", false, "bool flag")
    s := flag.String("s", "hello golang", "string flag")
    flag.Parse()
    fmt.Println("b is", *b)
    fmt.Println("s is", *s)
}
