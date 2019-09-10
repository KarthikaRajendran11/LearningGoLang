package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		io.WriteString(conn, "I see you connected")
		fmt.Fprintf(conn, "I see you connected")
		conn.Close()
	}
}