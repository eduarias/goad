package timer

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type inserter interface {
	calculateTokens(from, to time.Time) float64
	Delay() time.Duration
}

// Limiter ...
type Limiter struct {
	inserter inserter
	burst    int

	mu     sync.Mutex
	tokens float64
	// last is the last time the limiter's tokens field was updated
	last time.Time
}

// NewLimiter ...
func NewLimiter(b int, inserter inserter) *Limiter {
	return &Limiter{
		inserter: inserter,
		burst:    b,
	}
}

func (lim *Limiter) update() {
	lim.mu.Lock()
	lim.inserter.calculateTokens(lim.last, time.Now())
	lim.mu.Unlock()
}

// Wait ...
func (lim *Limiter) Wait(ctx context.Context) (err error) {
	return lim.WaitN(ctx, 1)
}

// WaitN ...
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error) {
	lim.mu.Lock()
	defer lim.mu.Unlock()

	if lim.burst <= n {
		return fmt.Errorf("rate: Wait(n=%d) would exceed context deadline", n)
	}

	lim.update()
	if float64(n) <= lim.tokens {
		lim.tokens -= float64(n)
		return nil
	}

	delay := lim.inserter.Delay()
	if delay == 0 {
		return nil
	}
	t := time.NewTimer(delay)
	defer t.Stop()
	select {
	case <-t.C:
		lim.update()
		lim.tokens -= float64(n)
		return nil
	case <-ctx.Done():
		// Context was canceled before we could proceed.  Cancel the
		// reservation, which may permit other events to proceed sooner.
		return ctx.Err()
	}

}
