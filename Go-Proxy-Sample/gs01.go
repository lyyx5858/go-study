package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	username, password := "foo", "bar"
	proxy := goproxy.NewProxyHttpServer()

	proxy.Verbose = true

	f1 := func(u, p string) bool { return u == username && p == password }
	ProxyBasicAuth(proxy, f1)

	//下面是原始代码，在上面两行，利用临时变量f1来代替下面这几行，充分理解函数就是一个变量的概念
	//ProxyBasicAuth(proxy, func(u, p string) bool {
	//	return u == username && p == password
	//})

	fmt.Println("start Server at port:8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}

//-------------------------------------------

var unauthorizedMsg = []byte("404 not found")

func BasicUnauthorized(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode:    404, // for purpose of avoiding proxy detection
		ProtoMajor:    1,
		ProtoMinor:    1,
		Request:       req,
		Header:        http.Header{},
		Body:          ioutil.NopCloser(bytes.NewBuffer(unauthorizedMsg)),
		ContentLength: int64(len(unauthorizedMsg)),
	}
}

func auth(req *http.Request, f func(user, passwd string) bool) bool {
	authheader := strings.SplitN(req.Header.Get(ProxyAuthHeader), " ", 2)
	req.Header.Del(ProxyAuthHeader)
	if len(authheader) != 2 || authheader[0] != "Basic" {
		return false
	}
	userpassraw, err := base64.StdEncoding.DecodeString(authheader[1])
	if err != nil {
		return false
	}
	userpass := strings.SplitN(string(userpassraw), ":", 2)
	if len(userpass) != 2 {
		return false
	}
	return f(userpass[0], userpass[1])
}

//==============================================================================================================
// Basic returns a basic HTTP authentication handler for requests
//
// You probably want to use auth.ProxyBasic(proxy) to enable authentication for all proxy activities

//func Basic(f func(user, passwd string) bool) goproxy.ReqHandler {
//	//FuncReqHandler是一种类型，下面这句话利用强制类型转换，将一个匿名函数转换成了FuncReqHandler类型
//	return goproxy.FuncReqHandler(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
//		if !auth(req, f) {
//			log.Fatal("basic auth verify for normal request failed")
//			return nil, BasicUnauthorized(req)
//		}
//		req.Header.Del(ProxyAuthHeader)
//		return req, nil
//	})
//}

//下面这个函数是上面的改写，只是为了让结构清楚。
func Basic(f func(user, passwd string) bool) goproxy.ReqHandler {
	temp := func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if !auth(req, f) {
			log.Fatal("basic auth verify for normal request failed")
			return nil, BasicUnauthorized(req)
		}
		req.Header.Del(ProxyAuthHeader)
		return req, nil
	}
	return goproxy.FuncReqHandler(temp) //此处的用法是强制类型转换，将temp转换成了FuncReqHandler
	// temp是个函数类型的变量，签名是：func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response)
	//强制类型转换的前提是：被转换的变量与转换类型签名一致。
}
//==============================================================================================================




// BasicConnect returns a basic HTTP authentication handler for CONNECT requests
//
// You probably want to use auth.ProxyBasic(proxy) to enable authentication for all proxy activities
func BasicConnect(f func(user, passwd string) bool) goproxy.HttpsHandler {
	return goproxy.FuncHttpsHandler(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		if !auth(ctx.Req, f) {
			log.Fatal("basic auth verify for connect request failed")
			ctx.Resp = BasicUnauthorized(ctx.Req)
			return goproxy.RejectConnect, host
		}
		ctx.Req.Header.Del(ProxyAuthHeader)
		return goproxy.OkConnect, host
	})
}

// ProxyBasic will force HTTP authentication before any request to the proxy is processed
func ProxyBasicAuth(proxy *goproxy.ProxyHttpServer, f func(user, passwd string) bool) {
	//原始语句如下，执行顺序是，从左向右
	//proxy.OnRequest().Do(Basic(f))
	temp := proxy.OnRequest()
	temp.Do(Basic(f))
	//此处，Basic(f)的返回值是一个Handler.

	proxy.OnRequest().HandleConnect(BasicConnect(f))
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
