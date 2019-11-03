package httpinstrumented

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func validateRequestMetric(t *testing.T, m *RequestMetrics) {
	assert.NotNil(t, m.StartTime)
	assert.InEpsilon(t, time.Now().Unix(), m.StartTime.Unix(), 1)
	hasDuration(t, m.TCPDialTime)
	hasDuration(t, m.RequestWriteTime)
	hasDuration(t, m.FirstResponseByteTime)
	assert.NotEmpty(t, m.HostPort)
}

func validateClientMetric(t *testing.T, m *ClientMetrics) {
	assert.NotNil(t, m.InitTime)
	assert.InDelta(t, time.Now().Unix(), m.InitTime.Unix(), 1)
	hasDuration(t, m.ElapsedTime)
}

func validateResponseMetric(t *testing.T, m *ResponseMetrics) {
	validateClientMetric(t, m.ClientMetrics)
	validateRequestMetric(t, m.RequestMetrics)
	assert.Greater(t, m.ElapsedTime.Nanoseconds(), m.FirstResponseByteTime.Nanoseconds())
}

func hasDuration(t *testing.T, duration time.Duration) {
	assert.NotNil(t, duration)
	assert.NotZero(t, duration.Nanoseconds())
	assert.InDelta(t, 0, duration.Milliseconds(), 100)
}
