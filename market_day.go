package marketday

import (
	"encoding/json"
	"time"
)

type Day struct {
	t time.Time
}

func NewDay(year int, month time.Month, day int) *Day {
	return &Day{time.Date(year, month, day, 15, 0, 0, 0, time.UTC)}
}

func NewDayFromTime(t time.Time) *Day {
	return NewDay(t.Year(), t.Month(), t.Day())
}

func (d *Day) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.t)
}

func (d *Day) UnmarshalJSON(data []byte) error {
	var t time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	d.t = t
	return nil
}

var holidays = map[time.Time]bool{
	NewDay(2023, 1, 2).t:   true,
	NewDay(2023, 1, 16).t:  true,
	NewDay(2023, 2, 20).t:  true,
	NewDay(2023, 4, 7).t:   true,
	NewDay(2023, 5, 29).t:  true,
	NewDay(2023, 6, 19).t:  true,
	NewDay(2023, 7, 4).t:   true,
	NewDay(2023, 9, 4).t:   true,
	NewDay(2023, 11, 23).t: true,
	NewDay(2023, 12, 25).t: true,
}

var halfDays = map[time.Time]bool{
	NewDay(2023, 7, 3).t:   true,
	NewDay(2023, 11, 24).t: true,
}

func (d *Day) Equal(d2 *Day) bool {
	return d.t.Equal(d2.t)
}

func (d *Day) IsZero() bool {
	return d.t.IsZero()
}

func (d *Day) IsMarketDay() bool {
	return d.isTradingDay() && !d.isHoliday()
}

func (d *Day) IsFullMarketDay() bool {
	return d.IsMarketDay() && !d.isHalfDay()
}

func (d *Day) previousDay() *Day {
	return &Day{d.t.AddDate(0, 0, -1)}
}

func (d *Day) PreviousMarketDay() *Day {
	day := d.previousDay()

	for !day.IsMarketDay() {
		day = day.previousDay()
	}

	return day
}

func (d *Day) PreviousMarketDays(n int) []*Day {
	days := make([]*Day, 0, n)
	day := d

	for len(days) < n {
		day = day.PreviousMarketDay()
		days = append(days, day)
	}

	return days
}

func (d *Day) isTradingDay() bool {
	switch d.t.Weekday() {
	case time.Saturday, time.Sunday:
		return false
	default:
		return true
	}
}

func (d *Day) isHoliday() bool {
	return holidays[d.t]
}

func (d *Day) isHalfDay() bool {
	return halfDays[d.t]
}
