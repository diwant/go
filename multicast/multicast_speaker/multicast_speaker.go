///////////////////
// Reference code at https://gist.github.com/fiorix/9664255
///////////////////

package main

import (
	"log"
	"net"
	"time"
)

const servAddr string = "224.0.0.1:9876"
const maxReadSize int = 8192

func main() {
	log.Println("Speakbot 2017 Activated.  Initializing...")
	log.Println("------------------------------------")
	ping(servAddr)
}

func ping(address string) {
	speakThis := "beep"

	log.Println("[x] Confirm UDP address resolvable, address is", address)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal("Could not resolve address", err)
	}

	log.Println("[x] Dialing UDP address", addr)
	c, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Could not Dial UDP address", err)
	}

	log.Println("[x] Connection success!")
	log.Println("[x] Commencing speech transmission")
	log.Println("------------------------------------")
	log.Println("")
	for {

		log.Println("Saying", speakThis)
		c.Write([]byte(speakThis))

		// Change to the next word to speak
		switch speakThis {
		case "beep":
			speakThis = "boop"
		case "boop":
			speakThis = "beep"
		}

		// Wait a second to say the next thing
		time.Sleep(1 * time.Second)
	}
}
