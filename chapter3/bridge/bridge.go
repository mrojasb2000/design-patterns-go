package main

import (
	"errors"
	"fmt"
	"io"
)

// acceptance criteria 1
type PrinterAPI interface {
	PrintMessage(string) error
}

// acceptance criteria 2
type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

// acceptance criteria 4
type PrinterAbstraction interface {
	Print() error
}

// acceptance criteria 5
type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

// acceptance criteria 6
type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PacktPrinter) Print() error {
	return errors.New("Not implemented yet")
}
