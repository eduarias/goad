package request

import "time"

// Metrics defines request measures
type Metrics struct {
	StartTime             time.Time
	DNSTime               time.Duration
	TCPDialTime           time.Duration
	TLSTime               time.Duration
	RequestWriteTime      time.Duration
	FirstResponseByteTime time.Duration
	HostPort              string
}
