package httpinstrumented

import "time"

// RequestMetrics defines request measures
type RequestMetrics struct {
	StartTime             time.Time
	DNSTime               time.Duration
	TCPDialTime           time.Duration
	TLSTime               time.Duration
	RequestWriteTime      time.Duration
	FirstResponseByteTime time.Duration
	HostPort              string
}

// ClientMetrics contains measures at client level
type ClientMetrics struct {
	InitTime    time.Time
	ElapsedTime time.Duration
}
