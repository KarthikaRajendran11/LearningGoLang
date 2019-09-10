package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	scanner := bufio.NewScanner(conn)
	data := make(map[string]string)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "Value : %s\n", v)
		case "SET":
			k := fs[1]
			v := fs[2]
			data[k] = v
			fmt.Fprintf(conn, "Key : %s, value : %s\n", k, v)
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Println(conn, "Invalid Command %s\n", fs[0])
			continue
		}
	}
}
