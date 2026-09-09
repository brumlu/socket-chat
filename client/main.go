package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"encoding/json"
)

func main () {

	serverConnection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer serverConnection.Close()

	fmt.Println("Enter your name to join the chat")
	
	var user string
	userNameInput := bufio.NewScanner(os.Stdin)
	if userNameInput.Scan() {
		user = userNameInput.Text()

	}


	go func() {
		serverScanner := bufio.NewScanner(serverConnection)
		for serverScanner.Scan() {
			text := serverScanner.Text()
			message, _ := FromJsonString(text)
			fmt.Printf("[%s]: %s\n", message.SenderName, message.MessageText)
		}

	if err := serverScanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	}()

	scannerInput := bufio.NewScanner(os.Stdin)
	fmt.Println("Connected to server. Type messages to send.")
	
	for {
		if !scannerInput.Scan() {
			break
		}

		msgText := scannerInput.Text()

		message := Message{
			MessageText: msgText,
			SenderName: user,
		}

		_, err := serverConnection.Write([]byte(message.ToJsonString()))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}

	if err := scannerInput.Err(); err != nil {
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