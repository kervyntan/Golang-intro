package main

import (
	"Golang-intro/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Kervyn")
	fmt.Println(message)
}
