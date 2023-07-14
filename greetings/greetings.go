package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// name string -> type of parameter name is string
// return type is string
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// Declaring and initializing message variable in one line
	message := fmt.Sprintf(RandomMessage(), name)
	return message, nil
}

func RandomMessage() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
