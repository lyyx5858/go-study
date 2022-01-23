package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"log"
	"net"
	"net/http"
)

var (
	listen  = flag.String("listen", "localhost:8080", "listen on address")
	ip      = flag.String("ip", "", "listen on address")
	verbose = flag.Bool("verbose", false, "verbose output")
)

func main() {
	flag.Parse()

	if *ip == "" {
		log.Fatal("IP address must be speicified")
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose
	proxy.Tr.Dial = func(network, addr string) (c net.Conn, err error) {
		if network == "tcp" {
			localAddr, err := net.ResolveTCPAddr(network, *ip+":0")
			if err != nil {
				return nil, err
			}
			remoteAddr, err := net.ResolveTCPAddr(network, addr)
			if err != nil {
				return nil, err
			}
			return net.DialTCP(network, localAddr, remoteAddr)
		}

		return net.Dial(network, addr)
	}
	log.Fatal(http.ListenAndServe(*listen, proxy))
}
