package greetings

import (
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Empty name")
	}

	return name, nil
}
