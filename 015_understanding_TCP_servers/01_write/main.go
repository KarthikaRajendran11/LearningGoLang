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

		io.WriteString(conn, "\nHello from TCP server")
		fmt.Fprintf(conn, "\nquestion 1")
		fmt.Fprintln(conn, "\n answer to question 1")

		conn.Close()
	}
}
