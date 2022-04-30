package main

import (
	"errors"
	"fmt"
)

func main() {
	// twentyFodderNoError says there are 20.0 fodder
	//fodder, err := DivideFood(ErrScaleMalfunction, 10)
	// fodder == 2.0
	// err == nil
	//fmt.Println("fodder: ", fodder, "err= ", err)

	// twentyFodderWithErrScaleMalfunction says there are 20.0 fodder and a ErrScaleMalfunction
	//fodder, err = DivideFood(twentyFodderWithErrScaleMalfunction, 10)
	// fodder == 4.0
	// err == nil
}

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	switch {
	case err == ErrScaleMalfunction && fodder > 0:
		fodder *= 2
	case fodder < 0 && (err == nil || err == ErrScaleMalfunction):
		return 0, errors.New("negative fodder")
	case err != nil:
		return 0, err
	case cows == 0:
		return 0, errors.New("division by zero")
	case cows < 0:
		return 0, &SillyNephewError{cows}
	}
	return fodder / float64(cows), nil

}

// WeightFodder returns the amount of available fodder.
type WeightFodder interface {
	FodderAmount() (float64, error)
}

// ErrScaleMalfunction indicates an error with the scale.
var ErrScaleMalfunction = errors.New("sensor error")

// See types.go for the types defined for this exercise.
// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}
