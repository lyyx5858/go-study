package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	http.ListenAndServe(":8080", mux)
}

type myHandler struct{} //此处myHandler是空结构 ，只是为了实现ServeHTTP方法，
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye World"))
}
