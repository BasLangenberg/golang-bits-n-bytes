package main

import (
	"errors"
	"fmt"
)

func GiveMeError(ok bool) (*string, error) {
	if ok {
		s := "ok"
		return &s, nil
	}

	return nil, fmt.Errorf("Tweede steen: %w", errors.New("Foutje, bedankt"))
}

func main() {
	_, err1 := GiveMeError(true)
	if err1 != nil {
		fmt.Println(errors.Unwrap(err1))
	} else {
		fmt.Println("Geen error 1")
	}

	_, err2 := GiveMeError(false)
	if err2 != nil {
		fmt.Println(errors.Unwrap(err2))
		fmt.Println(err2)
	} else {
		fmt.Println("Geen error 2")
	}
}
