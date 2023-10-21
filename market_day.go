package marketday

import (
	"time"
)

type Day struct {
	t time.Time
}

func MakeDay(year int, month time.Month, day int) Day {
	return Day{time.Date(year, month, day, 15, 0, 0, 0, time.UTC)}
}

var holidays = map[time.Time]bool{
	MakeDay(2023, 1, 2).t:   true,
	MakeDay(2023, 1, 16).t:  true,
	MakeDay(2023, 2, 20).t:  true,
	MakeDay(2023, 4, 7).t:   true,
	MakeDay(2023, 5, 29).t:  true,
	MakeDay(2023, 6, 19).t:  true,
	MakeDay(2023, 7, 4).t:   true,
	MakeDay(2023, 9, 4).t:   true,
	MakeDay(2023, 11, 23).t: true,
	MakeDay(2023, 12, 25).t: true,
}

var halfDays = map[time.Time]bool{
	MakeDay(2023, 7, 3).t:   true,
	MakeDay(2023, 11, 24).t: true,
}

func (d Day) IsMarketDay() bool {
	return d.isTradingDay() && !d.isHoliday()
}

func (d Day) IsFullMarketDay() bool {
	return d.IsMarketDay() && !d.isHalfDay()
}

func (d Day) PreviousDay() Day {
	return Day{d.t.AddDate(0, 0, -1)}
}

func (d Day) PreviousMarketDay() Day {
	day := d.PreviousDay()

	for !day.IsMarketDay() {
		day = day.PreviousDay()
	}

	return day
}

func (d Day) PreviousMarketDays(n int) []Day {
	days := make([]Day, 0, n)
	day := d

	for len(days) < n {
		day = day.PreviousMarketDay()
		days = append(days, day)
	}

	return days
}

func (d Day) isTradingDay() bool {
	switch d.t.Weekday() {
	case time.Saturday, time.Sunday:
		return false
	default:
		return true
	}
}

func (d Day) isHoliday() bool {
	return holidays[d.t]
}

func (d Day) isHalfDay() bool {
	return halfDays[d.t]
}
