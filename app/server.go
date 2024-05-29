package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	// Start the server and listen on port 4221
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221:", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a separate function
		go handleConnection(conn)
	}
}

// handleConnection handles incoming connections and sends the HTTP response
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Define the HTTP response
	response := "HTTP/1.1 200 OK\r\n\r\n"

	// Send the response to the client
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
	fmt.Println("Response sent:", response)
}
