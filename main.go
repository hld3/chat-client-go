package main

import (
	"log"
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
