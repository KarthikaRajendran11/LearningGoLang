package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"io"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			log.Fatalln("")
		}
		body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
		io.WriteString(conn, "\r\n")
	}
}