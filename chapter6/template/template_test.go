package template

import (
	"strings"
	"testing"
)

type TestStruct struct {
	Template
}

// Implicity implementing the MessageRetriever interface
func (m *TestStruct) Message() string {
	return "world"
}

func TestTemplate_ExecuteAlgorithm(t *testing.T) {

	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{}
		res := s.ExecuteAlgorithm(s)
		expected := "world"

		if !strings.Contains(res, expected) {
			t.Errorf("Expected string '%s' wasn't found on returned string\n", expected)
		}
	})
}
