package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

// Pointer to a struct of type singleton as nil
var instance *singleton

// Implement public method
func GetInstance() Singleton {
	if instance == nil {
		// Pointer to an instance of the type singleton
		instance = new(singleton)
	}
	return instance
}

// Implement interface singleton
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
