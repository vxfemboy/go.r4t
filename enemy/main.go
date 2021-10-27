package main

import "net"

func main { 
	sIP := "127.666.27.1"
	Port := "55555"
	sAddy := SIP + ":" + Port
	connc, err := net.Dial("tcp", sAddy)
	if err != nil{
		fmt.Println("Connection not established with server")
		log.Fatal(err)
	}
	fmt.Println("[+] Connection Established with: ", connec.RemoteAddr().String())
}