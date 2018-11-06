package main

import (
	"log"
	"testing"
)

func TestAthlete_Trait(t *testing.T) {
	athlete := Athlete{}
	athlete.Train()
}

func TestSwimmer_Swim(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	(swimmer.MySwim)()
}

func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}

func TestSwimmer_Swim2(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmer.Train()
	swimmer.Swim()
}

func TestTree(t *testing.T) {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
		},
		Left: &Tree{4, nil, nil},
	}

	println(root.Right.Right.LeafValue)
}

func TestSon_GetParentField(t *testing.T) {
	son := Son{}

	log.Println("son.P:", son.P)
	GetParentField(&son.P)
}
