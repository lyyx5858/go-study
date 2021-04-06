package main

import (
	// "github.com/elazarl/goproxy"
	"elazarl/goproxy"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	fmt.Print(proxy)
	proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*baidu.com$"))).
		HandleConnect(goproxy.AlwaysReject)

	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
