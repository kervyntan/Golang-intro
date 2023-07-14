package main

import (
	"Golang-intro/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	// message := greetings.Hello("Kervyn")
	fmt.Println(message)
}
