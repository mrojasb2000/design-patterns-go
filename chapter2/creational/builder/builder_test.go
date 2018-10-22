package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {

	const MOTORBIKE_SEATS_DEFAULT = 1
	const MOTORBIKE_WHEELS_DEFAULT = 2
	const CAR_WHEELS_DEFAULT = 4
	const BUS_WHEELS_DEFAULT = 8

	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != CAR_WHEELS_DEFAULT {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = MOTORBIKE_SEATS_DEFAULT

	if motorbike.Wheels != MOTORBIKE_WHEELS_DEFAULT {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n", motorbike.Structure)
	}

	busBuilder := &BusBuilder{}
	manufacturingComplex.SetBuilder(busBuilder)
	manufacturingComplex.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels != BUS_WHEELS_DEFAULT {
		t.Errorf("Wheels on a car must be 8 and they were %d\n", bus.Wheels)
	}

	if bus.Structure != "Bus" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", bus.Structure)
	}

	if bus.Seats != 30 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", bus.Seats)
	}
}
