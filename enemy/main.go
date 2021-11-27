package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	sIP := "127.0.0.1"
	Port := "5555"
	sAddy := sIP + ":" + Port
	connc, err := net.Dial("tcp", sAddy)
	if err != nil {
		fmt.Println("Connection not established with server")
		log.Fatal(err)
	}
	fmt.Println("[+] Connection Established with: ", connec.RemoteAddr().String())
	fmt.Println("brain melting commencing")
}
