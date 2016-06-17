package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var data = make(map[string]string)

func handle (conn net.Conn){
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		// Skip blank lines
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "GET":
			key := fields[1]
			value := data[key]
			fmt.Fprintf(conn, "%s\n", value)
		case "SET":
			if len(fields) != 3 {
				io.WriteString(conn, "EXPECTED VALUE\n")
				continue
			}
			key := fields[1]
			value := fields[2]
			data[key] = value
		case "DEL":
			key := fields[1]
			delete(data, key)
		default:
			io.WriteString(conn, "INVALID COMMAND " + fields[0] + "\n")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handle(connection)
	}
}