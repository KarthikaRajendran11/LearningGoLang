package main

import (
	"fmt"
	"bufio"
	"net"
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

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			panic(ln)
		}
		defer conn.Close()
		io.WriteString(conn, ln)
	}
	io.WriteString(conn, "Code got here")
}

func handle(conn net.Conn) {
	serve(conn)
}