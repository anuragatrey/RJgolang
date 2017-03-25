package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		handle(conn)
	}
}
