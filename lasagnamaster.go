package main

import "fmt"

func main() {

	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTimeB(layers, 3))
	// => 18
	fmt.Println(PreparationTimeB(layers, 0))
	// => 12

	fmt.Println(Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}))
	// => 100, 0.4

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}

	AddSecretIngredient(friendsList, myList)
	// myList => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"}
	fmt.Println(myList)

	quantities := []float64{1.2, 3.6, 10.5}
	fmt.Println(ScaleRecipe(quantities, 4))
	// => []float64{ 2.4, 7.2, 21 }
}

func PreparationTimeB(layers []string, prep int) int {
	if prep == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * prep
	}

}

func Quantities(layers []string) (nood int, sauce float64) {

	for _, v := range layers {
		if v == "noodles" {
			nood++
		} else {
			if v == "sauce" {
				sauce++
			}
		}
	}
	return nood * 50, sauce * 0.2
}

func AddSecretIngredient(fl []string, my []string) {
	my[len(my)-1] = fl[len(fl)-1]

}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaled := make([]float64, len(amounts))
	for i := range amounts {
		scaled[i] = amounts[i] * float64(portions) / 2
	}
	return scaled
}
