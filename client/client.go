package client

import (
	"encoding/binary"
	"fmt"
	"net"
	"tcp/utils"
)

type Client struct {
	conn net.Conn
	host string
	port int
}

// NewClient Creates a new client
func NewClient(host string, port int) *Client {
	return &Client{port: port, host: host}
}

// Close Disconnect current client
func (c *Client) Close() {
	err := c.conn.Close()

	if err != nil {
		return
	}
}

// Connect Creates a new connection
func (c *Client) Connect() {
	var err error

	c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port))

	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
}

// SendMessage send a new message to server
func (c *Client) SendMessage(message string) {
	encodedMessage := utils.EncodeMessage(message)

	// Send the size of the packet first
	sizeBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeBuf, uint32(len(encodedMessage)))
	c.conn.Write(sizeBuf)

	// Send the actual message (packet)
	c.conn.Write(encodedMessage)

	// Reading the response size
	responseSizeBuf := make([]byte, 4)

	var err error
	_, err = c.conn.Read(responseSizeBuf)

	if err != nil {
		fmt.Println("Error reading response size:", err)
		return
	}

	// Decode the response size
	responseSize := binary.BigEndian.Uint32(responseSizeBuf)

	// Read the actual response
	responseBuf := make([]byte, responseSize)
	_, err = c.conn.Read(responseBuf)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Decode and print the response
	decodedResponse := utils.DecodeMessage(responseBuf)
	fmt.Println("Server response:", decodedResponse)
}
