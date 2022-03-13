package main

import (
	"fmt"
)

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	TotalBirdCount(birdsPerDay)
	fmt.Println(TotalBirdCount(birdsPerDay))
	// => 34

	birdsPerDay = []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	BirdsInWeek(birdsPerDay, 2)
	fmt.Println(BirdsInWeek(birdsPerDay, 2))
	// => 12

	birdsPerDay = []int{2, 5, 0, 7, 4, 1}
	//FixBirdCountLog(birdsPerDay)
	fmt.Println(FixBirdCountLog(birdsPerDay))
	// => [3 5 1 7 5 1]
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var count int = 0
	for i := 0; i < len(birdsPerDay); i++ {
		count += birdsPerDay[i]
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	if week > 0 {
		return TotalBirdCount(birdsPerDay[week*7-7 : week*7])

	} else {
		return 0
	}
	//offset := 7 * (week - 1)
	//return TotalBirdCount(birdsPerDay[offset : offset+7])
	//return TotalBirdCount(birdsPerDay[(week-1)*7 : week*7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1 //birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
