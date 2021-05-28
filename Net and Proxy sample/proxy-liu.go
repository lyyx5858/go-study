package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	lPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 64*1024)
		},
	}
)

var i int

func main() {
	var port string = "8080"

	li, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println("error listen ", err)
		return
	}
	defer li.Close()

	log.Println("开启监听端口 " + port)

	var tempDelay time.Duration
	for {
		client, err := li.Accept() //此处阻塞
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Printf("server: Accept error: %v; retrying in %v\n", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return
		}
		tempDelay = 0

		go handle(client)
	}
}

func handle(client net.Conn) {
	i++
	fmt.Println("====================i=", i, "==============================")
	defer client.Close()

	req, err := http.ReadRequest(bufio.NewReader(client))
	if err != nil {
		log.Println("1:req read err:", err)
		return
	}
	defer req.Body.Close()

	host := req.Host
	if _, port, _ := net.SplitHostPort(host); port == "" {
		host = net.JoinHostPort(host, "443")
	}

	var server net.Conn

	retry:=3
	//server, err := net.Dial("tcp", address)
	for i := 0; i < retry; i++ {
		server, err = net.Dial("tcp", host)
		if err != nil {
			log.Println("2:Dial err:", err)
			time.Sleep(10*time.Millisecond)
		} else {
			break
		}

	}

	resp := &http.Response{
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     http.Header{},
	}
	resp.Header.Add("Proxy-Agent", "test")

	if err!= nil {
		resp.StatusCode = http.StatusServiceUnavailable
		resp.Write(client)
		return

	}

	defer server.Close()
	/*
		if method == "CONNECT" { //向请求方发出成功信号。HTTP 200
			_, err = fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
			if err != nil {
				log.Println(err)
				continue
			}

		} else {
			err = req.WriteProxy(server) //将client发来的请求，指向服务器
			if err != nil {
				log.Println(err)
				continue
			}
		}

		buf := make([]byte, 32*1024)

		err = req.Body.Close()
		if err != nil {
			log.Println(err)
		}
		go io.CopyBuffer(client, server, buf)
		go io.CopyBuffer(server, client, buf)
	*/

	fmt.Println(req.Method)
	if req.Method == http.MethodConnect {
		b := []byte("HTTP/1.1 200 Connection established\r\n" +
			"Proxy-Agent: test" + "\r\n\r\n")
		client.Write(b)
		if err != nil {
			log.Println("3:",err)
		}
	} else {
		req.Header.Del("Proxy-Connection")
		if err = req.WriteProxy(server); err != nil {
			log.Printf("4:[http] %s -> %s : %s\n", client.RemoteAddr(), client.LocalAddr(), err)
			return
		}
	}

	fmt.Printf("5:[http] %s <-> %s\n", client.RemoteAddr(), host)

	err = transport(client, server)
	if err != nil {
		log.Println("6:Copy Error:", err)
	}
	fmt.Printf("[http] %s >-< %s\n", client.RemoteAddr(), host)

}

func transport(rw1, rw2 io.ReadWriter) error {
	errc := make(chan error, 1)

	go func() {
		errc <- copyBuffer(rw1, rw2)
	}()

	go func() {
		errc <- copyBuffer(rw2, rw1)
	}()

	err := <-errc

	if err != nil && err == io.EOF {
		err = nil
	}
	return err
}

func copyBuffer(dst io.Writer, src io.Reader) error {
	buf := lPool.Get().([]byte)
	defer lPool.Put(buf)

	_, err := io.CopyBuffer(dst, src, buf)
	return err
}
