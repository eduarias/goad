package timer

import "time"

// LinearRate ...
type LinearRate struct {
	// rate defines the number of operations per second
	rate float64
}

// NewLinearRate ...
func NewLinearRate(rate float64) *LinearRate {
	return &LinearRate{rate: rate}
}

func (lr *LinearRate) calculateTokens(from, to time.Time) float64 {
	elapsed := to.Sub(from)
	return lr.rate * elapsed.Seconds()
}

// Delay ...
func (lr *LinearRate) Delay() time.Duration {
	return time.Duration(0)
}
