package main

func main() {

	needLicense := NeedsLicense("car")
	println("car? ", needLicense)
	// Output: true

	needLicense = NeedsLicense("bike")
	println("bike? ", needLicense)
	// Output: false

	vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
	println(vehicle)
	// Output: "Toyota Corolla is clearly the better choice."

	println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
	// Output: "Volkswagen Beetle is clearly the better choice."

	println(CalculateResellPrice(1000, 1))
	// Output: 800

	println(CalculateResellPrice(1000, 5))
	// Output: 700

	println(CalculateResellPrice(1000, 15))
	// Output: 500
}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var needed bool = false
	if kind == "car" || kind == "truck" {
		needed = true
	}
	return needed
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	choosed := option2 + " is clearly the better choice."
	if option1 < option2 {
		choosed = option1 + " is clearly the better choice."
	}
	return choosed
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var finalPrice float64
	if age < 3 {
		finalPrice = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		finalPrice = originalPrice * 0.7
	} else {
		finalPrice = originalPrice * 0.5
	}
	return finalPrice
}
