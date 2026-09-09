package main

import (
	"fmt"
	"net"
	"bufio"
	"encoding/json"
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

		message, err := FromJsonString(clientMsg)
		if err != nil {
			return
		}

		fmt.Printf("Message from %s: %s\n", message.SenderName, message.MessageText)
		
		_, err = connection.Write([]byte(message.ToJsonString()))
		if err != nil {
			fmt.Println("Error sending response:", err)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

type Message struct {
	SenderName string
	MessageText string
}

func FromJsonString(data string) (Message, error){
	
	var message Message
	
	if len(data) <= 0 {
		return Message{}, fmt.Errorf("Empty data string")
	}

	err := json.Unmarshal([]byte(data), &message)
	if err != nil {
		return Message{}, fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	return message, nil

}

func (m Message) ToJsonString() string {
	data, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s\n", string(data))
}