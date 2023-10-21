package marketday

import (
	"testing"
)

func TestKnownMarketDays(t *testing.T) {
	days := []Day{
		MakeDay(2023, 1, 3),
		MakeDay(2023, 4, 18),
		MakeDay(2023, 6, 22),
		MakeDay(2023, 7, 3),
		MakeDay(2023, 10, 31),
		MakeDay(2023, 11, 24),
	}

	for _, day := range days {
		if !day.IsMarketDay() {
			t.Fatalf("Expected %s to be a market day.", day.t)
		}

	}
}

func TestKnownHalfMarketDays(t *testing.T) {
	days := []Day{
		MakeDay(2023, 7, 3),
		MakeDay(2023, 11, 24),
	}

	for _, day := range days {
		if day.IsFullMarketDay() {
			t.Fatalf("Did not expect %s to be a full market day.", day.t)
		}
	}
}

func TestKnownNonMarketDays(t *testing.T) {
	days := []Day{
		MakeDay(2023, 1, 2),
		MakeDay(2023, 2, 18),
		MakeDay(2023, 4, 7),
		MakeDay(2023, 10, 29),
		MakeDay(2023, 11, 23),
	}

	for _, day := range days {
		if day.IsMarketDay() {
			t.Fatalf("Did not expect %s to be a market day.", day.t)
		}

	}
}

func TestPreviousMarketDay(t *testing.T) {
	cases := []struct {
		start    Day
		expected Day
	}{
		{MakeDay(2023, 1, 3), MakeDay(2022, 12, 30)},
		{MakeDay(2023, 4, 18), MakeDay(2023, 4, 17)},
		{MakeDay(2023, 1, 16), MakeDay(2023, 1, 13)},
		{MakeDay(2023, 11, 24), MakeDay(2023, 11, 22)},
		{MakeDay(2023, 7, 12), MakeDay(2023, 7, 11)},
	}

	for _, c := range cases {
		actual := c.start.PreviousMarketDay()
		if actual != c.expected {
			t.Fatalf("Expected previous market day %s, got %s", c.expected.t, actual.t)
		}
	}
}

func TestPreviousMarketDays(t *testing.T) {
	cases := []struct {
		start    Day
		expected []Day
	}{
		{MakeDay(2023, 1, 3), []Day{MakeDay(2022, 12, 30), MakeDay(2022, 12, 29)}},
		{MakeDay(2023, 4, 18), []Day{MakeDay(2023, 4, 17)}},
		{MakeDay(2023, 7, 1), []Day{
			MakeDay(2023, 6, 30),
			MakeDay(2023, 6, 29),
			MakeDay(2023, 6, 28),
			MakeDay(2023, 6, 27),
			MakeDay(2023, 6, 26),
			MakeDay(2023, 6, 23),
		}},
	}

	for _, c := range cases {
		actual := c.start.PreviousMarketDays(len(c.expected))
		aLen := len(actual)
		eLen := len(c.expected)

		if aLen != eLen {
			t.Fatalf("Expected %d previous market days, got %d", eLen, aLen)
		}

		for i, a := range actual {
			if a != c.expected[i] {
				t.Fatalf("Expected previous market day[%d] %s, got %s", i, c.expected[i].t, a.t)
			}
		}
	}
}
