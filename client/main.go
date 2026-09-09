package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main () {

	serverConnection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer serverConnection.Close()

/* 	fmt.Println("Enter your name to join the chat")
	
	var user string
	userNameInput := bufio.NewScanner(os.Stdin)
	if userNameInput.Scan() {
		user = userNameInput.Text()

	} */


	go func() {
		serverScanner := bufio.NewScanner(serverConnection)
		for serverScanner.Scan() {
			fmt.Println(serverScanner.Text())
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

		_, err := serverConnection.Write([]byte(msgText + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}

	if err := scannerInput.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}