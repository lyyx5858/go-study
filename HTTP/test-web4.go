package main

import (
	"fmt"
	"net/http"
)

type hotdog int

type test interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	var d hotdog
	var t test
	http.ListenAndServe(":8080", t)
	http.ListenAndServe(":9090", d)

}

func (t test)ServeHTTP(w http.ResponseWriter, r *http.Request) {




}


