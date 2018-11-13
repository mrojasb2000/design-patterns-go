package main

import (
	"errors"
	"strings"
	"testing"
)

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Content received on Writer was empty")
	return
}

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error tryping to use the API implementation: Message: %s\n", err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}

	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrorMessage := "You need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct.\n Actual: %s\nExpected: %s\n", err.Error(), expectedErrorMessage)
		}
	}
}
