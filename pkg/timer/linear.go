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
func (lr *LinearRate) Delay(neededTokens float64) time.Duration {
	if neededTokens <= 0 {
		return time.Duration(0)
	}
	return time.Duration(neededTokens/lr.rate) * time.Second
}
