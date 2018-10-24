package abstract_factory

type FamiliarCar struct{}

func (*FamiliarCar) GetDoors() int {
	return 5
}

func (*FamiliarCar) GetWheels() int {
	return 4
}

func (*FamiliarCar) GetSeats() int {
	return 5
}
