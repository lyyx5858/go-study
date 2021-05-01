package main

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"net/url"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	fmt.Println("Start Client at port:7070...")
	proxy.Verbose = true

	proxy.Tr.Proxy = func(request *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8080")
	}

	log.Fatal(http.ListenAndServe(":7070", proxy))
}
