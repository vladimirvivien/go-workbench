package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type proxy struct {
	net,
	localAddr,
	remoteAddr string
}

func Proxy(net, addr string) *proxy {
	return &proxy{net: net, localAddr: addr}
}

func (p *proxy) To(addr string) *proxy {
	p.remoteAddr = addr
	return p
}

func (p *proxy) Start() error {
	l, err := net.Listen(p.net, p.localAddr)
	if err != nil {
		return err
	}
	// block for conn
	lconn, err := l.Accept()
	if err != nil {
		return err
	}
	defer lconn.Close()
	defer l.Close()

	// dial remote
	dialer := new(net.Dialer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	rconn, err := dialer.DialContext(ctx, p.net, p.remoteAddr)
	if cancel != nil {
		cancel()
	}
	if err != nil {
		return err
	}
	defer rconn.Close()

	switch p.net {
	case "tcp", "tcp4", "tcp6":
		if conn, ok := rconn.(*net.TCPConn); ok {
			conn.SetKeepAlive(true)
			conn.SetKeepAlivePeriod(time.Minute * 1)
		}
	}

	errCh := make(chan error)
	go p.copy(errCh, lconn, rconn)
	go p.copy(errCh, rconn, lconn)

	return <-errCh
}

func (p *proxy) copy(errCh chan<- error, dst, src net.Conn) {
	_, err := io.Copy(dst, src)
	errCh <- err
}

func main() {
	var network, fromAddr, toAddr string
	flag.StringVar(&network, "n", "tcp", "network protocol to use (default tcp)")
	flag.StringVar(&fromAddr, "f", ":9090", "local address to use for proxy")
	flag.StringVar(&toAddr, "r", "", "remote address to proxy")
	flag.Parse()

	if toAddr == "" {
		fmt.Println("remote address must be specified with -r")
		os.Exit(1)
	}

	prox := Proxy(network, fromAddr).To(toAddr)
	fmt.Printf("Proxying %s -> %s ...\n", fromAddr, toAddr)
	if err := prox.Start(); err != nil {
		fmt.Println("proxy failed:", err)
		os.Exit(1)
	}
	fmt.Println("Proxy done.")
}
