package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

func main() {
	var port string = "8888"
	li, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println("error listen ", err)
		return
	}
	defer li.Close()
	log.Println("开启监听端口 " + port)

	var i int
	clientData := make([]byte, 4096) //空的4K的MAP

	for { //死循环监听
		client, err := li.Accept() //此处阻塞
		if err != nil {
			log.Println("监听错误", err)
			break
		}

		i++
		fmt.Println("====================i=", i, "==============================")

		n, err := client.Read(clientData) // 此处读client发过来的http请求至clientData, n是表示读了多少个字节
		requestStrings := string(clientData[:n])
		// log.Println("CLIENT request =========: \n\n", requestStrings) //RequestStrings是http 标准包头

		var method, uri, host, address string
		//取第一行进行分析
		fmt.Sscanf(string(clientData[:bytes.IndexByte(clientData, '\n')]), "%s%s", &method, &uri)
		key := strings.Split(requestStrings, "\n") //将HTTP包头用换行分成多行数组

		for _, c := range key { //将host从包头里分离出来 ，去除后面端口号，如：www.baidu.com:443变成 www.baidu.com
			if strings.Index(c, "Host") != -1 {
				host = c[strings.Index(c, "Host:")+5:]
				host = strings.TrimSpace(host) //去除前后的空格
				host = strings.Trim(host, ":443")
				host = strings.Trim(host, ":80")
			}
		}

		fmt.Println()
		fmt.Println("Method:----->", method)
		fmt.Println("uri:-------->", uri)
		fmt.Println("Host ------->", host)

		//host利用上面for已经分离出来，URI是最终访问地址, 将address从地址里分出来 ，加上端口号

		if strings.Index(uri, "http:") > 0 {
			parseUrl, err := url.Parse(uri)
			fmt.Println("parseurl is:", parseUrl)
			if err != nil {
				log.Println("error", err)
			}
			address = parseUrl.Host
		} else if strings.Index(uri, ":443") > 0 {
			address = host + ":443"
			fmt.Println("2")
		} else {
			address = host + ":80"
			fmt.Println("3")
		}

		fmt.Println("Address is ------->", address)

		//获得了请求host和port，开始拨号进行tcp连接
		server, err := net.Dial("tcp", address)
		if err != nil {
			log.Println(err)
			continue
		}

		if method == "CONNECT" { //向请求方发出成功信号。HTTP 200
			fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
			fmt.Println("5")

		} else {
			server.Write(clientData[:n]) //将client发来的请求，指向服务器
			fmt.Println("6")
		}

		go io.Copy(client, server)
		go io.Copy(server, client)

	}
}
