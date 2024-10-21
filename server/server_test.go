package server

import (
	"encoding/binary"
	"net"
	"tcp/utils"
	"testing"
)

func TestServerHandleConnection(t *testing.T) {
	server := NewServer(8080)

	var clientConn net.Conn
	server.conn, clientConn = net.Pipe()

	defer server.conn.Close()
	defer clientConn.Close()

	go server.handleConnection()

	// Send message from the client side
	message := "Hello Server"
	encodedMessage := utils.EncodeMessage(message)

	// Write message size and message
	sizeBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeBuf, uint32(len(encodedMessage)))
	clientConn.Write(sizeBuf)
	clientConn.Write(encodedMessage)

	// Read response size
	responseSizeBuf := make([]byte, 4)
	_, err := clientConn.Read(responseSizeBuf)
	if err != nil {
		t.Fatalf("Error reading response size: %v", err)
	}

	// Read response
	responseSize := binary.BigEndian.Uint32(responseSizeBuf)
	responseBuf := make([]byte, responseSize)
	_, err = clientConn.Read(responseBuf)
	if err != nil {
		t.Fatalf("Error reading response: %v", err)
	}

	// Verify response
	expectedResponse := "ACK: " + message
	decodedResponse := utils.DecodeMessage(responseBuf)
	if decodedResponse != expectedResponse {
		t.Errorf("Expected '%s', got '%s'", expectedResponse, decodedResponse)
	}
}
