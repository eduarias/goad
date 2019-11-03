package httpinstrumented

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptrace"
	"time"
)

// Request defines a requests with performance metrics
type Request struct {
	*http.Request
	Metrics *RequestMetrics
}

// NewRequestWithContext returns an standard http request powered by metrics
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error) {
	var dnsStart, connDialStart, tlsStart time.Time
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	m := &RequestMetrics{}

	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			m.HostPort = hostPort
			m.StartTime = time.Now()
		},
		GotConn:              func(connInfo httptrace.GotConnInfo) { time.Since(m.StartTime) },
		DNSStart:             func(dsi httptrace.DNSStartInfo) { dnsStart = time.Now() },
		DNSDone:              func(ddi httptrace.DNSDoneInfo) { m.DNSTime = time.Since(dnsStart) },
		ConnectStart:         func(network, addr string) { connDialStart = time.Now() },
		ConnectDone:          func(network, addr string, err error) { m.TCPDialTime = time.Since(connDialStart) },
		TLSHandshakeStart:    func() { tlsStart = time.Now() },
		TLSHandshakeDone:     func(cs tls.ConnectionState, err error) { m.TLSTime = time.Since(tlsStart) },
		WroteRequest:         func(wri httptrace.WroteRequestInfo) { m.RequestWriteTime = time.Since(m.StartTime) },
		GotFirstResponseByte: func() { m.FirstResponseByteTime = time.Since(m.StartTime) },
	}
	req.Close = true // closes file descriptor but prevent TCP to reuse connection
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	return &Request{req, m}, err
}

// NewRequest wraps NewRequestWithContext using the background context.
func NewRequest(method, url string, body io.Reader) (*Request, error) {
	return NewRequestWithContext(context.Background(), method, url, body)
}
