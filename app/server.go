package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "0.0.0.0:6379"

	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	// let's hardcode the response for now
	conn.Write([]byte("+PONG\r\n"))
}
