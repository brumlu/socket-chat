package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"client/ui"

	tea "charm.land/bubbletea/v2"
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

	model := ui.InitModel(serverConnection, user)
	p := tea.NewProgram(model)


	go func() {
		serverScanner := bufio.NewScanner(serverConnection)
		for serverScanner.Scan() {
			text := serverScanner.Text()
			message, _ := ui.FromJsonString(text)

			p.Send(message)
		}

	if err := serverScanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	}()

	if _, err := p.Run(); err != nil {
	fmt.Println("Error starting UI:", err)
	return
	}
}