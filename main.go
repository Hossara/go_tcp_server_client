package main

import (
	"sync"
	"tcp/client"
	"tcp/server"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go server.NewServer(8080).Connect()

	cl := client.NewClient("localhost", 8080)
	defer cl.Close()
	cl.Connect()

	cl.SendMessage("Message1")
	cl.SendMessage("Message2")
	cl.SendMessage("Message3")
	cl.SendMessage("Message4")
	cl.SendMessage("Message5")

	wg.Wait()
}
