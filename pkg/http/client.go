package httpinstrumented

import (
	"net/http"
	"time"
)

// Client is an instrumented  version of http.client
type Client struct {
	http.Client
	Metrics *ClientMetrics
}

// DefaultClient replaces http.DefaultClient
var DefaultClient = &Client{http.Client{}, &ClientMetrics{}}

// Do behave like http.Client but getting performance metrics
func (c *Client) Do(req *Request) (*http.Response, error) {
	if c.Metrics == nil {
		c.Metrics = &ClientMetrics{}
	}
	c.Metrics.InitTime = time.Now()
	defer func() { c.Metrics.ElapsedTime = time.Since(c.Metrics.InitTime) }()
	return c.Client.Do(req.Request)
}

// Get replaces the http.Get
func Get(url string) (resp *http.Response, err error) {
	return DefaultClient.Get(url)
}
