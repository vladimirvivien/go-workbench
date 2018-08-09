package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/tcpproxy"
)

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

	fmt.Printf("Proxying %s -> %s ...\n", fromAddr, toAddr)

	var prox tcpproxy.Proxy
	prox.AddRoute(fromAddr, tcpproxy.To(toAddr))
	log.Fatal(prox.Run())
	fmt.Println("Proxy done.")
}
