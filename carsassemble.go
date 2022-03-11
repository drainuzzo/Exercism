package main

import "fmt"

func main() {
	fmt.Println("CalculateWorkingCarsPerHour:", CalculateWorkingCarsPerHour(1547, 90))
	// Output: 1392.3
	fmt.Println("CalculateCost:", CalculateCost(37))
	// Output: 355000

	fmt.Println("CalculateCost:", CalculateCost(21))
	// Output: 200000

	fmt.Println("CalculateWorkingCarsPerMinute:", CalculateWorkingCarsPerMinute(1105, 90))
	// Output: 16
}

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * successRate / 100 / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	resto := uint(carsCount) % 10
	ris := uint(carsCount) / 10
	return resto*10000 + ris*95000
}
