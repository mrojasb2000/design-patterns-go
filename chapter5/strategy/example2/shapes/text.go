package shapes

import (
	"../strategy"
)

type TextSquare struct {
	// Struct embedded
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
