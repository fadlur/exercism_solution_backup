package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    timeParse, _ := time.Parse(layout, date)
    return timeParse
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t, _ := time.Parse("January 2, 2006 15:04:05", date)
    return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	time, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
    hour := time.Hour()
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	time := Schedule(date)
    return fmt.Sprintf("You have an appointment on %s at %s.",
                      time.Format("Monday, January 2, 2006,"),
                      time.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().UTC().Year()
    currentMonth := time.September
    currentDay := 15
	return time.Date(
        currentYear,
        currentMonth,
        currentDay,
        0, 0, 0, 0,
        time.UTC,
    )
}
