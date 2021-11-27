package main

import (
	"fmt"
	"log"
	"hconnc"
)

func main() {
	sIP := "127.0.0.1"
	Port := "5555"

	connc, err := hconnc.cserver(sIP, Port)
	if err != nil {
		log.Fatal(err)
	}
	defer connc.Close()
	fmt.Println("Connnected with Server :", connc.RemoteAddr().string())

}
