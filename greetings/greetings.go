package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hellos(names []string) (map[string]string, error) {
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

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	var message string
	message = fmt.Sprintf(randomFormat(), name)
	// message = fmt.Sprintf(randomFormat())
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Great to see you, %v!",
		"Hail, %v. well we meet again!",
		"Hi, %v. Welcome!",
	}

	return formats[rand.Intn(len(formats))]
}
