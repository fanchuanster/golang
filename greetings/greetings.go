package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// as the func name begins with upper-case letter, it's exported and can be used by other packages that import this package
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
