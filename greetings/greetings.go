package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello1(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Hello2(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello3(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
func randomFormat() string {
	formats := []string{
		"Hi,%v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello3(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
