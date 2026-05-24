package main

import (
	"fmt"
	"time"
)

func main() {

	//Functional option pattern
	var options []ServerOption
	options = append(options, WithHost("example.com"), WithPort(6000), WithTimeout(60*time.Second))
	server := NewServer(options...)
	fmt.Printf("Server running on %s:%d with timeout %s\n", server.Host, server.Port, server.Timeout)

	//Builder pattern
	server1, err := NewServerBuilder().
		SetHost("example.com").
		SetPort(6000).
		SetTimeout(60 * time.Second).
		Build()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Server running on %s:%d with timeout %s\n", server1.Host, server1.Port, server1.Timeout)
}
