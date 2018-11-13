package main

import (
	"errors"
	"io"
)

// acceptance criteria 1
type PrinterAPI interface {
	PrintMessage(string) error
}

// acceptance criteria 2
type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	return errors.New("Not implemented yet")
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	return errors.New("Not implemented yet")
}
