package marketday

import (
	"math"
	"time"
)

type EODData struct {
	Symbol string
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

func sameFloat(a float64, b float64) bool {
	return math.Abs(a-b) <= 1e-9
}

func (a *EODData) Equal(b *EODData) bool {
	return a.Symbol == b.Symbol &&
		a.Date.Equal(b.Date) &&
		sameFloat(a.Open, b.Open) &&
		sameFloat(a.High, b.High) &&
		sameFloat(a.Low, b.Low) &&
		sameFloat(a.Close, b.Close) &&
		a.Volume == b.Volume
}
