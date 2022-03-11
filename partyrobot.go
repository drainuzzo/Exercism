package main

import (
	"fmt"
	"strconv"
)

func main() {
	println("Welcome:\n" + Welcome("Christiane"))
	// Output: Welcome to my party, Christiane!

	println("HappyBirthday:\n" + HappyBirthday("Frank", 58))
	// Output: Happy birthday Frank! You are now 58 years old!

	println("AssignTable:\n" + AssignTable("Christiane", 7, "Frank", "on the left", 23.7834298))
	// Output:
	// Welcome to my party, Christiane!
	// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
	// You will be sitting next to Frank.

}

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"

}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

//La seguente func sostituibile semplicemente con un Sprintf(table %03d) -.-
func GetDigits(table int) string {
	strtbl := strconv.Itoa(table)
	for i := 0; i < len(strtbl); i++ {
		if len(strtbl) < 3 {
			strtbl = "0" + strtbl
		}
	}
	return strtbl
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %s. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, GetDigits(table), direction, distance, neighbor)
}
