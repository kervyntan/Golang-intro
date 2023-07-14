package main

import (
	"Golang-intro/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// slice of names:
	names := []string{"Kervyn", "John", "Vinh"}

	// message, err := greetings.Hello("")
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
