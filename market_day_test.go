package marketday

import (
	"encoding/json"
	"testing"
)

type SampleConfig struct {
	CurrentDay *Day `json:"currentDay"`
}

func TestKnownMarketDays(t *testing.T) {
	days := []*Day{
		NewDay(2023, 1, 3),
		NewDay(2023, 4, 18),
		NewDay(2023, 6, 22),
		NewDay(2023, 7, 3),
		NewDay(2023, 10, 31),
		NewDay(2023, 11, 24),
	}

	for _, day := range days {
		if !day.IsMarketDay() {
			t.Fatalf("Expected %s to be a market day.", day.t)
		}

	}
}

func TestKnownHalfMarketDays(t *testing.T) {
	days := []*Day{
		NewDay(2023, 7, 3),
		NewDay(2023, 11, 24),
	}

	for _, day := range days {
		if day.IsFullMarketDay() {
			t.Fatalf("Did not expect %s to be a full market day.", day.t)
		}
	}
}

func TestKnownNonMarketDays(t *testing.T) {
	days := []*Day{
		NewDay(2023, 1, 2),
		NewDay(2023, 2, 18),
		NewDay(2023, 4, 7),
		NewDay(2023, 10, 29),
		NewDay(2023, 11, 23),
	}

	for _, day := range days {
		if day.IsMarketDay() {
			t.Fatalf("Did not expect %s to be a market day.", day.t)
		}

	}
}

func TestPreviousMarketDay(t *testing.T) {
	cases := []struct {
		start    *Day
		expected *Day
	}{
		{NewDay(2023, 1, 3), NewDay(2022, 12, 30)},
		{NewDay(2023, 4, 18), NewDay(2023, 4, 17)},
		{NewDay(2023, 1, 16), NewDay(2023, 1, 13)},
		{NewDay(2023, 11, 24), NewDay(2023, 11, 22)},
		{NewDay(2023, 7, 12), NewDay(2023, 7, 11)},
	}

	for _, c := range cases {
		actual := c.start.PreviousMarketDay()
		if !actual.Equal(c.expected) {
			t.Fatalf("Expected previous market day %s, got %s", c.expected.t, actual.t)
		}
	}
}

func TestPreviousMarketDays(t *testing.T) {
	cases := []struct {
		start    *Day
		expected []*Day
	}{
		{NewDay(2023, 1, 3), []*Day{NewDay(2022, 12, 30), NewDay(2022, 12, 29)}},
		{NewDay(2023, 4, 18), []*Day{NewDay(2023, 4, 17)}},
		{NewDay(2023, 7, 1), []*Day{
			NewDay(2023, 6, 30),
			NewDay(2023, 6, 29),
			NewDay(2023, 6, 28),
			NewDay(2023, 6, 27),
			NewDay(2023, 6, 26),
			NewDay(2023, 6, 23),
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
			if !a.Equal(c.expected[i]) {
				t.Fatalf("Expected previous market day[%d] %s, got %s", i, c.expected[i].t, a.t)
			}
		}
	}
}

func TestMarshalJSON(t *testing.T) {

}

func TestUnmarshalJSON(t *testing.T) {
	configJSON := []byte(`{"currentDay": "2023-10-18T15:00:00Z"}`)

	var config *SampleConfig
	if err := json.Unmarshal(configJSON, &config); err != nil {
		t.Fatalf("Failed to unmarshal json: %v", err)
	}

	expected := NewDay(2023, 10, 18)
	if !config.CurrentDay.Equal(expected) {
		t.Fatalf("Expected %s after unmarshal, got %s", expected.t, config.CurrentDay.t)
	}
}
