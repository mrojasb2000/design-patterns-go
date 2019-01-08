package memento

import (
	"fmt"
)

// Core business object
type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{}
}

func (o *originator) ExtractAndStoreState(m memento) {
	//Does nothing
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	//Does nothing
}

func (c *careTaker) Memento(i int) (memento, error) {
	return memento{}, fmt.Errorf("Not implemented yet")
}
