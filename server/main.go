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
	
	connc, err := hconnc.cvictim(IP, Port)
	if err != nil{ 
		log.Fatal(err)
	}
	defer connc.Close()
	fmt.Println("[+] Connected with ", connc.RemoteAddr().String())
}
