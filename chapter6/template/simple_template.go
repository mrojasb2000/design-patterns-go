package template

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
	return ""
}

func (t *Template) third() string {
	return ""
}

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return ""
}
