package main

import (
	"errors"
)

func main() {
	err := doesReturnError()
	if err != nil {
		panic(err)
	}
}

func doesReturnError() error {
	err := errors.New("this function simply returns an error")
	return err
}
