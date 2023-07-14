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

	// Range of random integer within the length of the formats slice
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	// Syntax:
	// make(map[key-type]value-type)
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}
