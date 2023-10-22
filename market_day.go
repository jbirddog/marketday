package marketday

import (
	"time"
)

var holidays = map[time.Time]bool{
	Day(2023, 1, 2):   true,
	Day(2023, 1, 16):  true,
	Day(2023, 2, 20):  true,
	Day(2023, 4, 7):   true,
	Day(2023, 5, 29):  true,
	Day(2023, 6, 19):  true,
	Day(2023, 7, 4):   true,
	Day(2023, 9, 4):   true,
	Day(2023, 11, 23): true,
	Day(2023, 12, 25): true,
}

var halfDays = map[time.Time]bool{
	Day(2023, 7, 3):   true,
	Day(2023, 11, 24): true,
}

func Day(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 15, 0, 0, 0, time.UTC)
}

func PreviousMarketDay(date time.Time) time.Time {
	date = date.AddDate(0, 0, -1)

	for !IsMarketDay(date) {
		date = date.AddDate(0, 0, -1)
	}

	return date
}

func PreviousMarketDays(date time.Time, count int) []time.Time {
	days := make([]time.Time, count, count)

	for i := 0; i < count; i++ {
		date = PreviousMarketDay(date)
		days[i] = date
	}

	return days
}

func IsMarketDay(date time.Time) bool {
	return isTradingDay(date) && !isHoliday(date)
}

func IsFullMarketDay(date time.Time) bool {
	return IsMarketDay(date) && !isHalfDay(date)
}

func isTradingDay(date time.Time) bool {
	switch date.Weekday() {
	case time.Saturday, time.Sunday:
		return false
	default:
		return true
	}
}

func isHoliday(date time.Time) bool {
	return holidays[date]
}

func isHalfDay(date time.Time) bool {
	return halfDays[date]
}
