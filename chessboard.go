//package chessboard
package main

import "fmt"

func main() {

	// boardvar := map[string][]bool{
	// 	"A": {true, false, true, false, true, false, false, false},
	// 	"B": {false, false, true, false, false, false, false, true},
	// 	"C": {false, true, true, false, false, true, false, false},
	// }

	board := Chessboard{
		"A": {true, false, true, false, true, false, false, false},
		"B": {false, false, true, false, false, false, false, true},
		"C": {false, true, true, false, false, true, false, false},
		"D": {true, true, true, false, false, true, false, false},
		"E": {false, false, false, false, false, false, false, false},
		"F": {false, true, true, false, false, true, false, false},
		"G": {true, true, true, false, false, true, false, false},
		"H": {false, true, true, false, false, true, false, true},
	}
	fmt.Println(CountInRank(board, "A"))
	// => 3
	fmt.Println(CountInFile(board, 2))
	// => 5
	fmt.Println(CountAll(board))
	// => 64
	fmt.Println(CountOccupied(board))
	// => 23

}

// Declare a type named Rank which stores if a square is occupied
// by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight
// Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	count := 0
	v, ok := cb[rank]
	if !ok {
		return count
	}
	for _, value := range v {
		if value {
			count++
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0
	if file < 0 || file > 8 {
		return count
	}
	for _, v := range cb {
		if v[file-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for range cb {
		count++
	}
	return count * count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, value := range cb {
		for _, v := range value {
			if v {
				count++
			}
		}
	}
	return count
}
