package main

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	fmt.Println("start Server at port:8080")

	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
