package request

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptrace"
	"time"
)

// Request defines a requests with performance metrics
type Request struct {
	*http.Request
	Metrics *Metrics
}

// NewRequest returns a new request structure with performance metrics
func NewRequest(method, url string, body io.Reader) (*Request, error) {
	var dnsStart, connDialStart, tlsStart time.Time
	req, err := http.NewRequest(method, url, body)
	m := &Metrics{}

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
