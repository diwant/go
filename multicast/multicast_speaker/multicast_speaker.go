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
