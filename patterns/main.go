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
	pizza, err := NewPizzaBuilder().
		SetDough("stuffed crust").
		SetSauce("marinara").
		SetSize("Large").
		Build()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Pizza order: %+v\n", pizza)
}
