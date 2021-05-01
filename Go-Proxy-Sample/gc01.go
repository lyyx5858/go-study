package main

import (
	"encoding/base64"
	"fmt"
	goproxy "github.com/elazarl/goproxy"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	username, password := "foo", "bar"

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	proxy.Tr.Proxy = func(request *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8080")
	}

	// proxy.ConnectDial = proxy.NewConnectDialToProxy(proxyUrl)
	proxy.ConnectDial = proxy.NewConnectDialToProxyWithHandler("http://127.0.0.1:8080",
		SetAuthForBasicConnectRequest(username, password))

	// set basic auth
	proxy.OnRequest().Do(SetAuthForBasicRequest(username, password))
	//上面这句话是在一定条件下，做。。。。
	//OnRequest()是条件，它本质是一个proxy的方法，其参数是可变，也可省略，省略代表任何条件下都要执行Do方法
	//Do的输入参数是一个handler，然后Do将这个新的handler附加到proxy的handler数组上去


	fmt.Println("Start Client at port:7070...")
	log.Fatal(http.ListenAndServe(":7070", proxy))
}

func SetAuthForBasicRequest(username, password string) goproxy.ReqHandler {
	return goproxy.FuncReqHandler(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		SetBasicAuth(username, password, req)
		return req, nil
	})
}

func SetAuthForBasicConnectRequest(username, password string) func(req *http.Request) {
	return func(req *http.Request) {
		SetBasicAuth(username, password, req)
	}
}

//COMMON
const (
	ProxyAuthHeader = "Quic-Proxy-Authorization"
)

func SetBasicAuth(username, password string, req *http.Request) {
	req.Header.Set(ProxyAuthHeader, fmt.Sprintf("Basic %s", basicAuth(username, password)))
}

func basicAuth(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func GetBasicAuth(req *http.Request) (username, password string, ok bool) {
	auth := req.Header.Get(ProxyAuthHeader)
	if auth == "" {
		return
	}

	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
