package main

import (
	"fmt"
	"time"
)

// Command patterns
type Command interface {
	Info() string
}

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

// Command handler
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

	timeCommand := &TimePassed{start: time.Now()}
	first.Next(timeCommand)
	messageCommand := &CustomMessage{message: "Hello"}
	first.Next(messageCommand)
}
