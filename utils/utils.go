package utils

// DecodeMessage Decode function to process received data
func DecodeMessage(data []byte) string {
	// Returns a UTF-8 encoded string
	return string(data)
}

// EncodeMessage Encode function to encode response
func EncodeMessage(response string) []byte {
	return []byte(response)
}
