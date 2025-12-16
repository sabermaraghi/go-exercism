package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layouts := []string{
        "1/2/2006 15:04:05",
        "January 2, 2006 15:04:05",
        "Monday, January 2, 2006 15:04:05",
    }

    for _, layout := range layouts {
        t, err := time.Parse(layout, date)
        if err == nil {
            return t
        }
    }
    return time.Time{}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    appointment := Schedule(date)

    if appointment.IsZero() {
        return false
    }
    return time.Now().After(appointment)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointment := Schedule(date)
    return appointment.Hour() >= 12 && appointment.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointment := Schedule(date)
    formt := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%02d.", appointment.Weekday(), appointment.Month(), appointment.Day(), appointment.Year(), appointment.Hour(), appointment.Minute())
    return formt
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
