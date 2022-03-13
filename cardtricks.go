package main

import "fmt"

func main() {

	card, ok := GetItem([]int{1, 2, 4, 1}, 2)
	fmt.Println(card)
	// Output: 4
	fmt.Println(ok)
	// Output: true

	cardD, okK := GetItem([]int{1, 2, 4, 1}, 10)
	fmt.Println(cardD)
	// Output: 0
	fmt.Println(okK)
	// Output: false

	index := 2
	newCard := 6
	fmt.Println(SetItem([]int{1, 2, 4, 1}, index, newCard))
	// Output: []int{1, 2, 6, 1}

	indexX := -1
	newCardD := 6
	fmt.Println(SetItem([]int{1, 2, 4, 1}, indexX, newCardD))
	// Output: []int{1, 2, 4, 1, 6}

	fmt.Println(PrefilledSlice(8, 3))
	// Output: []int{8, 8, 8}

	fmt.Println(RemoveItem([]int{3, 2, 6, 4, 8}, 1))
	//Output: []int{3, 2, 4, 8}

	fmt.Println(RemoveItem([]int{3, 2, 6, 4, 8}, 11))
	//Output: []int{3, 2, 6, 4, 8}

}

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index < 0 || index > len(slice)-1 {
		return 0, false
	}
	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice

}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length > 0 {
		slice := make([]int, length)
		for i := 0; i < length; i++ {
			slice[i] = value
		}
		return slice
	} else {
		return []int{}
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var sliceN = []int{}
	if index == len(slice)-1 {
		return slice[:index]
	} else if index >= 0 && index < len(slice) {
		for i := index; i < len(slice)-1; i++ {
			sliceN = append(slice[:i], slice[i+1])
		}
	} else {
		return slice
	}
	return sliceN

	//migliore
	// if index < 0 || index > len(slice)-1 {
	// 	return slice
	// }
	// return append(slice[:index], slice[index+1:]...)
}
