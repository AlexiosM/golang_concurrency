package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// TODO: write server program to handle concurrent client connections.
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(connection)
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			log.Fatal(err)
			continue
		}
		time.Sleep(time.Second)
	}
}
