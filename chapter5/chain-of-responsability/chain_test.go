package chain_of_responsability

import (
	"strings"
	"testing"
)

type myTestWriter struct {
	receivedMessage *string
}

// Implements the common Write method from the io.Writer interface
func (m *myTestWriter) Write(p []byte) (int, error) {
	tempMessage := string(p)
	m.receivedMessage = &tempMessage
	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}

func TestCreateDefaultChain(t *testing.T) {
	//Our test ChainLogger
	myWriter := myTestWriter{}

	writerLogger := WriterLogger{Writer: &myWriter}
	second := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of them writes to console, second only if it founds "+
		"the world 'hello', thrid writes to some variable if second found 'hello'", func(t *testing.T) {

		chain.Next("message that breaks the chain")

		if myWriter.receivedMessage != nil {
			t.Error("Last link should not receive any message")
		}

		chain.Next("Hello")

		if myWriter.receivedMessage == nil || !strings.Contains(*myWriter.receivedMessage, "Hello") {
			t.Fatal("Last link didn't received expected message")
		}
	})
}
