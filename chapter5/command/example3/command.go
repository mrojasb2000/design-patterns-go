package main

import (
	"fmt"
	"time"
)

// Command patterns
type Command interface {
	Info() string
}

// Command handler
type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type CustomMessage struct {
	message string
}

func (c *CustomMessage) Info() string {
	return c.message
}

// Chain of responsability patterns
type ChainLogger interface {
	Next(Command)
}

type Logger struct {
	NextChain ChainLogger
}

func (l *Logger) Next(c Command) {
	time.Sleep(time.Second)

	fmt.Printf("Elapsed time from creation: %s\n", c.Info())

	if l.NextChain != nil {
		l.NextChain.Next(c)
	}
}

func main() {
	last := new(Logger)
	third := Logger{NextChain: last}
	second := Logger{NextChain: &third}
	first := Logger{NextChain: &second}

	command1 := &TimePassed{start: time.Now()}
	first.Next(command1)
	command := &CustomMessage{message: "Hello"}
	first.Next(command)
}
