package main

import (
	"fmt"
	"net/http"
)

type H1 struct{} //H1是个结构 ，但它实现了方法ServeHTTP
func (h H1) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello 1 Handler!")
}

func hello2 (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello 2!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello1 := H1{}

	http.Handle("/hello1", hello1) //Hello1 是个接口，此接口必须 实现ServHTTP方法
	http.HandleFunc("/hello2", hello2) // hello2是个函数，可以不用实现ServHTTP方法，因此HandleFunc这个函数会将其转换为
											// 和Hello1一样的接口类型，并拥有ServHTTP方法。
	server.ListenAndServe()
}