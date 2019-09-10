package main

import (
	"fmt"
	"net"
	"io"
	"bufio"
	"log"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "I see you connected")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if(ln == ""){
			panic("")
		}
		fmt.Println(ln)
	}
	defer conn.Close()
}