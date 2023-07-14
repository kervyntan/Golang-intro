package greetings

import (
	"errors"
	"fmt"
)

// name string -> type of parameter name is string
// return type is string
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// Declaring and initializing message variable in one line
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
