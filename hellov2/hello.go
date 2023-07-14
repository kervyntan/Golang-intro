package main

import (
	"Golang-intro/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("")
	message, err := greetings.Hello("Kervyn")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
