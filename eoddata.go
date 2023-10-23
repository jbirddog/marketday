package marketday

import (
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
