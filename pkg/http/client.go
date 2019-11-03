package httpinstrumented

import (
	"io"
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
func (c *Client) Do(req *Request) (*Response, error) {
	if c.Metrics == nil {
		c.Metrics = &ClientMetrics{}
	}
	c.Metrics.InitTime = time.Now()
	httpResp, err := c.Client.Do(req.Request)
	c.Metrics.ElapsedTime = time.Since(c.Metrics.InitTime)
	if err != nil {
		return nil, err
	}
	return &Response{*httpResp, &ResponseMetrics{req.Metrics, c.Metrics}}, err
}

// Get replaces the http.Client Get
func (c *Client) Get(url string) (resp *Response, err error) {
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Get replaces the http.Get
func Get(url string) (resp *Response, err error) {
	return DefaultClient.Get(url)
}

// Post replaces the http.Client Post
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error) {
	req, err := NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

// Post replaces the http.Post
func Post(url, contentType string, body io.Reader) (resp *Response, err error) {
	return DefaultClient.Post(url, contentType, body)
}
