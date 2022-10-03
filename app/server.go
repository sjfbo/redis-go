package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	addr := "0.0.0.0:6379"

	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer l.Close()

	conn, err := l.Accept()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		len, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			return
		}

		fmt.Println("Received data: ", string(buf[:len]))

		if strings.Contains(string(buf[:len]), "ping") {
			conn.Write([]byte("+PONG\r\n"))
		}
	}
}
