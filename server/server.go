package server

import (
	"encoding/binary"
	"fmt"
	"net"
	"tcp/utils"
)

type Server struct {
	port int
	conn net.Conn
}

func NewServer(port int) *Server {
	return &Server{port: port}
}

// Close Disconnect current client
func (c *Server) Close() {
	err := c.conn.Close()

	if err != nil {
		return
	}
}

func (s *Server) handleConnection() {
	defer s.conn.Close()

	for {
		// Reading the size of the incoming packet
		sizeBuf := make([]byte, 4)
		_, err := s.conn.Read(sizeBuf)
		if err != nil {
			fmt.Println("Error reading size:", err)
			return
		}

		// Decode the size of the packet
		packetSize := binary.BigEndian.Uint32(sizeBuf)

		// Reading the packet itself
		packetBuf := make([]byte, packetSize)
		_, err = s.conn.Read(packetBuf)
		if err != nil {
			fmt.Println("Error reading packet:", err)
			return
		}

		// Decode and process the message
		decodedMessage := utils.DecodeMessage(packetBuf)
		fmt.Println("Received:", decodedMessage)

		// Prepare and send response
		response := "ACK: " + decodedMessage
		encodedResponse := utils.EncodeMessage(response)

		// Send response size first
		responseSize := make([]byte, 4)
		binary.BigEndian.PutUint32(responseSize, uint32(len(encodedResponse)))
		s.conn.Write(responseSize)

		// Send the actual response
		s.conn.Write(encodedResponse)
	}
}

func (s *Server) Connect() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")
	for {
		var err error
		s.conn, err = listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go s.handleConnection()
	}
}
