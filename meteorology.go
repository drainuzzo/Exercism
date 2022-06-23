package main

import "fmt"

func main() {

	//temperatureUnit := Celsius
	celsiusUnit := Celsius
	fmt.Sprint(celsiusUnit)
	celsiusUnit.String()
	// => °C

	fahrenheitUnit := Fahrenheit
	fmt.Sprint(fahrenheitUnit)
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

	mphUnit := MilesPerHour
	mphUnit.String()
	// => mph
	fmt.Sprint(mphUnit)
	// => mph

	kmhUnit := KmPerHour
	kmhUnit.String()
	// => km/h
	fmt.Sprint(kmhUnit)
	// => km/h

	windSpeedYesterday := Speed{
		magnitude: 22,
		unit:      MilesPerHour,
	}
	windSpeedYesterday.String()
	// => 22 mph
	fmt.Sprint(windSpeedYesterday)
	// => 22 mph

	windSpeedNow := Speed{
		magnitude: 18,
		unit:      KmPerHour,
	}
	windSpeedNow.String()
	// => 18 km/h
	//fmt.Sprintf(windSpeedNow)
	// => 18 km/h

	sfData := MeteorologyData{
		location: "San Francisco",
		temperature: Temperature{
			degree: 57,
			unit:   Fahrenheit,
		},
		windDirection: "NW",
		windSpeed: Speed{
			magnitude: 19,
			unit:      MilesPerHour,
		},
		humidity: 60,
	}

	sfData.String()
	// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
	fmt.Sprint(sfData)
	// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity

}

type SpeedUnit int

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

func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

func (su SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	//fmt.Println(units)
	return units[su]
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (md MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", md.location, md.temperature, md.windDirection, md.windSpeed, md.humidity)
}

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (tu TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	//fmt.Println(units)
	return units[tu]
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}
