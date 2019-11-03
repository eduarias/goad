package httpinstrumented

import "net/http"

// Response adds metrics to http.Request
type Response struct {
	http.Response
	Metrics *ResponseMetrics
}
