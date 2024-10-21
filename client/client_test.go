package client

import (
	"encoding/binary"
	"net"
	"tcp/utils"
	"testing"
)

func TestClientSendMessage(t *testing.T) {
	// Use net.Pipe to simulate a client-server connection
	serverConn, clientConn := net.Pipe()
	defer serverConn.Close()
	defer clientConn.Close()

	// Create a mock client with the simulated connection
	mockClient := &Client{
		conn: clientConn,
		host: "localhost",
		port: 8080,
	}

	// Simulate server reading the message
	go func() {
		// Read message size
		sizeBuf := make([]byte, 4)
		_, err := serverConn.Read(sizeBuf)
		if err != nil {
			t.Fatalf("Error reading size: %v", err)
		}

		// Read message
		packetSize := binary.BigEndian.Uint32(sizeBuf)
		packetBuf := make([]byte, packetSize)
		_, err = serverConn.Read(packetBuf)
		if err != nil {
			t.Fatalf("Error reading message: %v", err)
		}

		// Assert received message
		expectedMessage := "TestMessage"
		decodedMessage := utils.DecodeMessage(packetBuf)
		if decodedMessage != expectedMessage {
			t.Errorf("Expected '%s', got '%s'", expectedMessage, decodedMessage)
		}

		// Send a response back to the client
		response := "ACK: " + decodedMessage
		encodedResponse := utils.EncodeMessage(response)

		// Write response size and response
		responseSize := make([]byte, 4)
		binary.BigEndian.PutUint32(responseSize, uint32(len(encodedResponse)))
		serverConn.Write(responseSize)
		serverConn.Write(encodedResponse)
	}()

	// Test sending a message from the client
	mockClient.SendMessage("TestMessage")
}
