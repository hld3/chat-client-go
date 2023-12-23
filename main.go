package main

import (
	"fmt"
	"log"
	"net"
)

// define the server address and port
const addr = "localhost"
const port = "8081"

func main() {
	err := StartServer(addr, port)
	if err != nil {
		log.Println("Error starting the server.", err)
	}

}

func handleRequest(conn net.Conn) {
	// TODO move to a separate package.
	fmt.Println("Confirmed Request Recieved, closing.")
	conn.Close()
}
