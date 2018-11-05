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

func main() {
	/*
		swimmer := CompositeSwimmerA{
			MySwim: Swim,
		}

		// Zero initilization struct
		swimmer.MyAthlete.Train()
		swimmer.MySwim()
	*/

	/*
		// A pointer to function
		localSwim := Swim

		swimmer := CompositeSwimmerA{
			MySwim: localSwim,
		}

		swimmer.MyAthlete.Train()
		swimmer.MySwim()

		// Embed objects
		fish := Shark{
			Swim: Swim,
		}

		fish.Eat()
		fish.Swim()
	*/

	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmer.Train()
	swimmer.Swim()
}
