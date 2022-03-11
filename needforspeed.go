package main

import "fmt"

func main() {

	speed := 5
	batteryDrain := 2

	//distance default struct value = 0
	car := NewCar(speed, batteryDrain)
	fmt.Println(car)
	// Output: Car{speed: 5, batteryDrain: 2, battery:100, distance: 0}

	distance := 800
	raceTrack := NewTrack(distance)
	fmt.Println(raceTrack)
	// Output: Track{distance: 800}

	speed = 5
	batteryDrain = 2
	car = Drive(car)
	fmt.Println(car)
	// Output: Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}

	car = NewCar(speed, batteryDrain)
	distance = 251
	raceTrack = NewTrack(distance)
	//car = dist:100 speed:5 battery:98 bd: 2
	fmt.Printf("distance: %d, car: %+v\n", distance, car)
	CanFinish(car, raceTrack)
	fmt.Println(CanFinish(car, raceTrack))
	// Output: true

}

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	modCar := Car{100, batteryDrain, speed, 0}
	return modCar
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	newTrack := Track{distance}
	return newTrack
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery > 0 && car.battery > car.batteryDrain {
		battery := 100 - car.batteryDrain
		drive := Car{car.speed, car.batteryDrain, battery, car.distance + car.speed}
		return drive
	} else {
		return car
	}

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	maxDistance := car.battery * car.speed / car.batteryDrain
	if track.distance <= maxDistance {
		return true
	}
	return false
}
