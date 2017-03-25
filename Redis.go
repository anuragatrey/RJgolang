package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

var data = make(map[string]string)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		//skip blank lines
		if len(fs) < 2 {
			continue
		}

		switch fs[0] {
		case "GET", "Get", "get":
			key := fs[1]
			value := data[key]
			fmt.Fprintf(conn, "%s\n", value)

		case "SET", "Set", "set":
			if len(fs) < 3 {
				io.WriteString(conn, "Expected value \n")
				continue
			}
			key := fs[1]
			value := fs[2]
			data[key] = value

		case "DEL":
			key := fs[1]
			delete(data, key)

		default:
			io.WriteString(conn, "Invalid Command"+fs[0]+"\n")
		}
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
