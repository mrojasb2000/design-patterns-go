package main

import (
	"fmt"
)

type RailroadWideChecker interface {
	CheckRailsWidth() int
}

// Train station
type Railroad struct {
	Width int
}

// Check whether a train fits the needs of the railroad
func (r *Railroad) IsCorrectSizeTrain(rwch RailroadWideChecker) bool {
	return rwch.CheckRailsWidth() != r.Width
}

// Train implements RailroadWideChecker interface
type Train struct {
	TrainWidth int
}

// Implement interface
func (t *Train) CheckRailsWidth() int {
	return t.TrainWidth
}

func main() {
	// Railsroad object
	railroad := Railroad{Width: 10}

	// Train objects
	passengerTrain := Train{TrainWidth: 10}
	cargoTrain := Train{TrainWidth: 15}

	canPassengerTrainPass := railroad.IsCorrectSizeTrain(&passengerTrain)
	canCargoTrainPass := railroad.IsCorrectSizeTrain(&cargoTrain)

	fmt.Println("Can passenger train pass?", canPassengerTrainPass)
	fmt.Println("Can cargo train pass?", canCargoTrainPass)
}
