package template

import (
	"strings"
)

type MessageRetriever interface {
	Message() string
}

type Templater interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type Template struct{}

func (t *Template) first() string {
	return "hello"
}

func (t *Template) third() string {
	return "template"
}

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}
