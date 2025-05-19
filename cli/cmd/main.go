package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting memo-box-kv ....")
	listener, err := net.Listen("tcp", ":8080") // Bind to port 8080
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080")

	for {
		conn, err := listener.Accept() // Wait for client
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		go handleConnection(conn) // Handle in new goroutine
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		input := string(buffer[:n])
		fmt.Println("Received:", input)

		conn.Write([]byte("You said: " + input))
	}
}
