package main

import (
	"strconv"
	"time"
)

func main() {

	Schedule("7/25/2019 13:45:00")
	// Output: 2019-07-25 13:45:00 +0000 UTC

	HasPassed("July 25, 2019 13:45:00")
	// Output: true

	IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
	// Output: true

	Description("7/25/2019 13:45:00")
	// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."

	AnniversaryDate()
	// Output: 2020-09-15 00:00:00 +0000 UTC

}

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	//standard layout Mon Jan 2 15:04:05 -0700 MST 2006
	layout := "1/02/2006 15:04:05"
	//date = "7/25/2019 13:45:00"
	t, _ := time.Parse(layout, date)
	//fmt.Println("t: ", t) //OUTPUT ok is: 2019-07-25 13:45:00 +0000 UTC
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	//ex: July 25, 2019 13:45:00
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	//fmt.Println(t.Before(time.Now()))
	return t.Before(time.Now())

}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	//Thursday, July 25, 2019 13:45:00
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	/* if t.Hour() >= 12 && t.Hour() < 18 {
		fmt.Println("true")
		return true
	} */
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	//Description("7/25/2019 13:45:00")
	// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
	layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	s := t.Format("Monday, January 2, 2006, at 15:04")
	return "You have an appointment on " + s + "."
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	curYear := time.Now().Year()
	date := strconv.Itoa(curYear) + "-09-15 00:00:00 +0000 UTC"
	layout := "2006-01-2 15:04:05 -0700 MST"
	t, _ := time.Parse(layout, date)
	//fmt.Println(t)
	return t
}
