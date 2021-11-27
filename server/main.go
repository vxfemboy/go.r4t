package main

import (
	"fmt"
	"log"
	"net"
)

//main
func main() {
	var connc net.Conn
	IP := "127.0.0.1"
	Port := "5555"
	Addy := IP + ":" + Port

	lis, err := net.Listen("tcp", Addy)
	if err != nil {
		log.Fatal(err)
	}
	connc, err = lis.Accept()
	if err != nil {
		log.Println("[-] Connection not established!")
	} else {
		fmt.Println("[+] Connection established from: " + connc.RemoteAddr().String())
		fmt.Println("brain melting commencing")
	}
}
