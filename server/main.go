package main

import (
	"fmt"
	"net"
	"bufio"
)

func main () {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		go handleConnection(connection)
		}
	}

func handleConnection(connection net.Conn) {
	
	defer connection.Close()

	fmt.Println("New client connected")

	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		clientMsg := scanner.Text()
		fmt.Println("Received from client:", clientMsg)

		response := fmt.Sprintf("Server received: %s\n", clientMsg)
		
		_, err := connection.Write([]byte(response))
		if err != nil {
			fmt.Println("Error sending response:", err)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}