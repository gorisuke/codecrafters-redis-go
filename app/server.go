package main

import (
	"fmt"
	"net"
	"os"
)

const pongResponse = "+PONG\r\n"

func handler(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		conn.Read(buf)
		conn.Write([]byte(pongResponse))
	}

}
func main() {
	fmt.Println("Start Processing")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handler(conn)
	}
}
