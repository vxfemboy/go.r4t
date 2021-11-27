package hconnc

import (
	"net"
)

func cserver(sIP string, Port string) (connc net.Conn, err error) {
	sAddy := sIP + ":" + Port
	connc, err := net.Dial("tcp", sAddy)
	if err != nil {
		return
	}
	return
}
