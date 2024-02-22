package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func numPack(packNum int) ([]byte, error) {
	list := []byte{}

	if packNum > 65535 {
		return list, errors.New("number > 65535")
	}

	list = append(list, uint8(packNum/256))
	list = append(list, uint8(packNum%256))

	return list, nil
}

func numUnpack(packData []byte) int {
	return (int(packData[0])*256) + int(packData[1])
}

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
	conn.Write([]byte("Hello, world!"))
	defer conn.Close()
}
