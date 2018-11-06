package main

import (
	"fmt"
)

// Form: Direct composite
type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming!")
}

//--------------------------------------------------------------------------
// Form: Embed composite
type Animal struct{}

func (r *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	//Embed object
	Animal
	Swim func()
}

//--------------------------------------------------------------------------
// Form: Interfaces
type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

//--------------------------------------------------------------------------
// Form: Binary Tree
type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

//--------------------------------------------------------------------------
// Form: Son
type Parent struct {
	SomeField int
}

type Son struct {
	//Parent    // Embed version - error
	P Parent // Without embed version - ok
}

func GetParentField(p *Parent) int {
	fmt.Println(p.SomeField)
	return 0
}
