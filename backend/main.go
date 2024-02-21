package main

import "fmt"
import "net"
import "os"

func main() {
	ln, err := net.Listen("tcp", ":8000")
	fmt.Println("Listening on port :8000")

	if err != nil {
		fmt.Println("Failed to listen on port :8000! Aborting...")
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		fmt.Println("Recieved connection")

		if err != nil {
			fmt.Println("Failed to fetch connection! Ignoring...")
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
}
