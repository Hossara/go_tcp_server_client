package utils

import (
	"testing"
)

func TestEncodeMessage(t *testing.T) {
	message := "Hello"
	encoded := EncodeMessage(message)
	if string(encoded) != message {
		t.Errorf("Expected '%s', got '%s'", message, string(encoded))
	}
}

func TestDecodeMessage(t *testing.T) {
	encoded := []byte("Hello")
	decoded := DecodeMessage(encoded)
	if decoded != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", decoded)
	}
}
