///////////////////
// Reference code at https://gist.github.com/fiorix/9664255
///////////////////

package main

import (
	"encoding/hex"
	"log"
	"net"
	"time"
)

const servAddr string = "224.0.0.1:9876"
const maxReadSize int = 8192

func main() {
	go ping(servAddr)
	listenForMulticast(servAddr, castHandler)
}

func ping(address string) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal("Could not resolve address", err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	for {
		c.Write([]byte("hum"))
		time.Sleep(1 * time.Second)
	}
}

func castHandler(a *net.UDPAddr, i int, b []byte) {
	log.Println(i, "bytes read from", a)
	log.Println(hex.Dump(b[:i]))
}

func listenForMulticast(address string, handle func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal("Could not resolve address", err)
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	l.SetReadBuffer(maxReadSize)
	for {
		b := make([]byte, maxReadSize)
		i, srcAddr, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("Read From Multicast failed", err)
		}
		handle(srcAddr, i, b)
	}
}
