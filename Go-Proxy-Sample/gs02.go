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

	ProxyBasic(proxy, "CSI",f1)  //这个函数的目的是将认证过程注册成为Proxy结构体的ReqHandler数组上去.
								//所谓Handler，其实就是一个接口。
								//所谓接口，就是定义了一个或者多个方法。
								//所谓方法，就是一个包含宿主（接收者）的函数
								//所谓函数就是一小块内存，包含可执行代码



	//下面是原始代码，在上面两行，利用临时变量f1来代替下面这几行，充分理解函数就是一个变量的概念
	//ProxyBasicAuth(proxy, func(u, p string) bool {
	//	return u == username && p == password
	//})

	fmt.Println("start Server at port:8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}


var unauthorizedMsg1 = []byte("407 Proxy Authentication Required")

var proxyAuthorizationHeader = "Proxy-Authorization"



// ProxyBasic will force HTTP authentication before any request to the proxy is processed
func ProxyBasic(proxy *goproxy.ProxyHttpServer, realm string, f func(user, passwd string) bool) {
	proxy.OnRequest().Do(Basic1(realm, f))
	proxy.OnRequest().HandleConnect(BasicConnect1(realm, f))
}

// Basic returns a basic HTTP authentication handler for requests
//
// You probably want to use auth.ProxyBasic(proxy) to enable authentication for all proxy activities
func Basic1(realm string, f func(user, passwd string) bool) goproxy.ReqHandler {
	return goproxy.FuncReqHandler(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if !auth1(req, f) {
			return nil, BasicUnauthorized1(req, realm)
		}
		return req, nil
	})
}

// BasicConnect returns a basic HTTP authentication handler for CONNECT requests
//
// You probably want to use auth.ProxyBasic(proxy) to enable authentication for all proxy activities
func BasicConnect1(realm string, f func(user, passwd string) bool) goproxy.HttpsHandler {
	return goproxy.FuncHttpsHandler(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		if !auth1(ctx.Req, f) {
			ctx.Resp = BasicUnauthorized1(ctx.Req, realm)
			return goproxy.RejectConnect, host
		}
		return goproxy.OkConnect, host
	})
}


func auth1(req *http.Request, f func(user, passwd string) bool) bool {
	authheader := strings.SplitN(req.Header.Get(proxyAuthorizationHeader), " ", 2)
	req.Header.Del(proxyAuthorizationHeader)
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
	fmt.Println(userpass[0],userpass[1])
	return f(userpass[0], userpass[1])
}

func BasicUnauthorized1(req *http.Request, realm string) *http.Response {
	// TODO(elazar): verify realm is well formed
	return &http.Response{
		StatusCode: 407,   //407就是让浏览器出现提示框，让用户手工输用户名和密码
		ProtoMajor: 1,
		ProtoMinor: 1,
		Request:    req,
		Header: http.Header{
			"Proxy-Authenticate": []string{"Basic realm=" + realm},
			"Proxy-Connection":   []string{"close"},
		},
		Body:          ioutil.NopCloser(bytes.NewBuffer(unauthorizedMsg1)),
		ContentLength: int64(len(unauthorizedMsg1)),
	}
}