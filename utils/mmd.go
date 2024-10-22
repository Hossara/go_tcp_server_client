package utils

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City  string `json:"city"`
	State string `json:"state,omitempty"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

func main() {
	// Example JSON with nested structure
	jsonData := `{
		"name": "John Doe",
		"age": 30,
		"email": "john@example.com",
		"address": {
			"city": "New York",
			"state": "NY"
		}
	}`

	data := map[string]interface{}{
		"name":  "John Doe",
		"age":   30,
		"email": "john@example.com",
	}

	jsonData2, _ := json.Marshal(data)

	println(string(jsonData2))

	// Declare a variable of type Person
	var person Person

	// Decode (Unmarshal) the JSON data into the Go struct
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Access the decoded nested data
	fmt.Printf("Name: %s, Age: %d, City: %s, State: %s\n", person.Name, person.Age, person.Address.City, person.Address.State)
}
