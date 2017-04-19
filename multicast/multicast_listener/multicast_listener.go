///////////////////
// Reference code at https://gist.github.com/fiorix/9664255
///////////////////

package main

import (
	"encoding/hex"
	"log"
	"net"
)

const servAddr string = "224.0.0.1:9876"
const maxReadSize int = 8192

func main() {
	log.Println("Welcome to Multicast Listen-o-Tron 2.0")
	log.Println("------------------------------------")
	listenForMulticast(servAddr, castHandler)
}

func castHandler(a *net.UDPAddr, i int, b []byte) {
	log.Println(i, "bytes read from", a)
	log.Println(hex.Dump(b[:i]))
}

func listenForMulticast(address string, handle func(*net.UDPAddr, int, []byte)) {

	log.Println("...resolving address for listening on UDP at", address)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal("Could not resolve address", err)
	}

	log.Println("...initializing listen for multicast udp at", addr)
	l, err := net.ListenMulticastUDP("udp", nil, addr)

	log.Println("...setting read buffer max to", maxReadSize)
	l.SetReadBuffer(maxReadSize)

	log.Println("...setup complete!")

	log.Println("...listening...")
	log.Println("------------------------------------")
	log.Println("")
	for {
		b := make([]byte, maxReadSize)
		i, srcAddr, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("Read From Multicast failed", err)
		}
		handle(srcAddr, i, b)
	}
}
