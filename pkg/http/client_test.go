package httpinstrumented

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoInstrumentationHTTPS(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	}))
	defer ts.Close()

	req, _ := NewRequest("GET", ts.URL, nil)
	client := Client{*ts.Client(), nil}
	resp, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDoInstrumentationHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	}))
	defer ts.Close()

	req, _ := NewRequest("GET", ts.URL, nil)
	client := Client{}
	resp, err := client.Do(req)

	m := client.Metrics
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotNil(t, m.InitTime)
	assert.InDelta(t, time.Now().Unix(), m.InitTime.Unix(), 1)
	hasDuration(t, m.ElapsedTime)
}

func TestGetInstrumentation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	}))
	defer ts.Close()

	_, _ = Get(ts.URL)

	// TODO - Get metrics for this request, right now has no asserts
}
