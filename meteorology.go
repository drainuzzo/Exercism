package main

import "fmt"

func main() {

	//temperatureUnit := Celsius
	celsiusUnit := Celsius
	fahrenheitUnit := Fahrenheit

	celsiusUnit.String()
	// => °C
	fahrenheitUnit.String()
	// => °F
	fmt.Sprint(celsiusUnit)
	// => °C

	fmt.Println("test")

	celsiusTemp := Temperature{
		degree: 21,
		unit:   Celsius,
	}
	celsiusTemp.String()
	// => 21 °C
	fmt.Sprint(celsiusTemp)
	// => 21 °C

	fahrenheitTemp := Temperature{
		degree: 75,
		unit:   Fahrenheit,
	}
	fahrenheitTemp.String()
	// => 75 °F
	fmt.Sprint(fahrenheitTemp)
	// => 75 °F

}

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
/* func (d Distance) String() string {
	return fmt.Sprintf("%v %v", d.number, d.unit)
} */

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (tt TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[tt]
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

// Add a String method to the Temperature type

/* type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData */
