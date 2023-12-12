package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Not Found Error")
)

func getById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "jhon" {
		return NotFoundError
	}
	return nil
}

func main() {
	i := getById("jhon")
	if i != nil {
		if errors.Is(i, ValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(i, NotFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Error tidak diketahui")
		}

	}
}
