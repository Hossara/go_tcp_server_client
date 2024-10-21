package main

import (
	"fmt"
	"sync"
	"tcp/client"
	"tcp/server"
)

func main() {
	// Initialize a wait group
	wg := sync.WaitGroup{}

	// Wait for 1 server instance
	wg.Add(1)

	// Create and connect server instance (listen on localhost port 8080)
	srv := server.NewServer(8080)
	go srv.Connect()

	// Create new client instance and connect to server
	cl := client.NewClient("localhost", 8080)

	// Make sure both server connection will be closed at the end
	defer srv.Close()

	// Connect to client
	cl.Connect()

	for {
		fmt.Print("Type your message: ")
		var message string

		fmt.Scanln(&message)

		if message != "end!" {
			cl.SendMessage(message)
		} else {
			cl.Close()
			wg.Done()
			break
		}
	}

	wg.Wait()

	fmt.Println("Bye Bye!")
}
