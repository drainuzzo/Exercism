package main

import "fmt"

func main() {
	// TODO: define the 'OvenTime' constant
	//const OvenTime = 40
	PreparationTime(2)
	ElapsedTime(2, 30)
	RemainingOvenTime(30)

}

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	a := OvenTime - actualMinutesInOven
	//fmt.Println(a)
	return a
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	tot := 2 * numberOfLayers
	return tot
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	end := PreparationTime(numberOfLayers) + actualMinutesInOven
	fmt.Println("end:", end)
	return end

}
