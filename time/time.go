// Package ntpclient allows go programs for retrieving the current time
// from an NTP server. This is a client only. The NTP version supported
// is version 4. Also compatible with sntp.
// Inspired by: https://github.com/lettier/ntpclient
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

// GetNetworkTime retrieves the current UTC time from the remote NTP server
// It returns the go time.Time type of the current time in UTC.
func main() {
	var host string
	flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
	flag.Parse()

	packet := make([]byte, 48)
	packet[0] = 0x1B

	addr, _ := net.ResolveUDPAddr("udp4", host)
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(10 * time.Second))

	_, err = conn.Write(packet)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Read(packet)
	if err != nil {
		log.Fatal(err)
	}

	//retrieve the bytes that we need for the current timestamp
	//data format is unsigned 64 bit long, big endian order
	//see: http://play.golang.org/p/6KRE-2Hq6n
	second := uint64(packet[40])<<24 | uint64(packet[41])<<16 | uint64(packet[42])<<8 | uint64(packet[43])
	fraction := uint64(packet[44])<<24 | uint64(packet[45])<<16 | uint64(packet[46])<<8 | uint64(packet[47])

	fmt.Printf("seconds=%0b [%d]\n", second, second)
	nsec := (second * 1e9)
	fmt.Printf("nanoseconds=%0b [%d]\n", nsec, nsec)

	fmt.Printf("fraction=%0b [%d]\n", fraction, fraction)
	nf := ((fraction * 1e9) >> 32)
	fmt.Printf("nanofracs=%0b [%d]\n", nf, nf)

	ntotal := nsec + nf
	fmt.Printf("ntotal=%0b [%d]\n", ntotal, ntotal)

	now := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(ntotal))
	fmt.Println(now)
}
