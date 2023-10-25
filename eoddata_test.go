package marketday

import (
	"testing"
)

func TestSameEODData(t *testing.T) {
	eod1 := &EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00}
	eod2 := &EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00}

	if !eod1.Equal(eod1) {
		t.Fatal("Expected eod1 to be equal")
	}

	if !eod1.Equal(eod2) {
		t.Fatal("Expected eods to be equal")
	}
}

func TestDifferentEODData(t *testing.T) {
	cases := []struct {
		a *EODData
		b *EODData
	}{
		{
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00},
			&EODData{"DEF", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00},
		},
		{
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00},
			&EODData{"ABC", Day(2023, 10, 26), 1.00, 2.00, 3.00, 4.00, 5.00},
		},
		{
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00},
			&EODData{"ABC", Day(2023, 10, 25), 1.01, 2.00, 3.00, 4.00, 5.00},
		},
		{
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00},
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.01, 3.00, 4.00, 5.00},
		},
		{
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00},
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.01, 4.00, 5.00},
		},
		{
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00},
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.01, 5.00},
		},
		{
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.00},
			&EODData{"ABC", Day(2023, 10, 25), 1.00, 2.00, 3.00, 4.00, 5.01},
		},
	}

	for i, c := range cases {
		if c.a.Equal(c.b) {
			t.Fatalf("Expected eods[%d] to not be equal", i)
		}
	}
}
