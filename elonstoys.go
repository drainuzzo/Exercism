package main

import (
	"fmt"
)

func main() {

	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	car.Drive()
	fmt.Println("car is now ", car)
	// car is now Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}

	fmt.Println(car.DisplayDistance())
	// Output: "Driven 0 meters"

	fmt.Println(car.DisplayBattery())
	// Output: "Battery at 100%"

	trackDistance := 100

	fmt.Println(car.CanFinish(trackDistance))
	// Output: true
}

//package elon

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
	maxDistance := c.battery * c.speed / c.batteryDrain
	if trackDistance <= maxDistance {
		return true
	}
	return false
}
