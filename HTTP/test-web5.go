package main

import (
	"fmt"
	"net/http"
)

//Handler嵌套

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello!")

}
func log(h http.HandlerFunc) http.HandlerFunc { //注意http.HandlerFuc 与 http.HandleFunc的区别
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Again!\n")
		h(w, r)
	}
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()

}
