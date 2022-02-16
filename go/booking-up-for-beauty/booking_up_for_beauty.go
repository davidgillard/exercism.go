// file: booking_up_for_beauty.go
// author: David Gillard
// date: 16/02/2022
// http://patorjk.com/software/taag

/*
  ____              _    _                              __             ____                   _
 |  _ \            | |  (_)                            / _|           |  _ \                 | |
 | |_) | ___   ___ | | ___ _ __   __ _   _   _ _ __   | |_ ___  _ __  | |_) | ___  __ _ _   _| |_ _   _
 |  _ < / _ \ / _ \| |/ / | '_ \ / _` | | | | | '_ \  |  _/ _ \| '__| |  _ < / _ \/ _` | | | | __| | | |
 | |_) | (_) | (_) |   <| | | | | (_| | | |_| | |_) | | || (_) | |    | |_) |  __/ (_| | |_| | |_| |_| |
 |____/ \___/ \___/|_|\_\_|_| |_|\__, |  \__,_| .__/  |_| \___/|_|    |____/ \___|\__,_|\__,_|\__|\__, |
                                  __/ |       | |                                                  __/ |
                                 |___/        |_|                                                 |___/
*/

package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
// Description("7/25/2019 13:45:00")
// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Schedule(date string) time.Time {
	parseDate, _ := time.Parse("1/2/2006 15:04:05", date)
	return parseDate
}

// HasPassed returns whether a date has passed
// HasPassed("July 25, 2019 13:45:00")
// Output: true
func HasPassed(date string) bool {
	parseDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(parseDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
// IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
// Output: true
func IsAfternoonAppointment(date string) bool {
	parseDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return parseDate.Hour() >= 12 && parseDate.Hour() < 18
}

// Description returns a formatted string of the appointment time
// Description("7/25/2019 13:45:00")
// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
// like the format is the same that the Schedule function
func Description(date string) string {
	parseDate := Schedule(date).Format("Monday, January 2, 2006, at 15:04.")
	return fmt.Sprintf("You have an appointment on %s", parseDate)
}

// AnniversaryDate returns a Time with this year's anniversary
// AnniversaryDate()
// Output: 2020-09-15 00:00:00 +0000 UTC
func AnniversaryDate() time.Time {
	parseDate, _ := time.Parse("2006-01-2", fmt.Sprintf("%v-09-15", time.Now().Year()))
	return parseDate
}
