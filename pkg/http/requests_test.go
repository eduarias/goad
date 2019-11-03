package httpinstrumented

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	req, _ := NewRequest("GET", ts.URL, nil)
	client := ts.Client()
	client.Do(req.Request)

	m := req.Metrics
	assert.NotNil(t, req.Metrics.StartTime)
	assert.InEpsilon(t, time.Now().Unix(), m.StartTime.Unix(), 1)
	hasDuration(t, m.TCPDialTime)
	hasDuration(t, m.TLSTime)
	hasDuration(t, m.RequestWriteTime)
	hasDuration(t, m.FirstResponseByteTime)
	assert.NotEmpty(t, m.HostPort)
}
