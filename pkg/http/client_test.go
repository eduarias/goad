package httpinstrumented

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	validateClientMetric(t, client.Metrics)
}

func TestDoInstrumentationError(t *testing.T) {
	req, err := NewRequest("GET", "http://localhost:9999", nil)
	client := Client{}
	_, err = client.Do(req)
	assert.Error(t, err)
}

func TestGetInstrumentation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	}))
	defer ts.Close()

	resp, err := Get(ts.URL)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	validateResponseMetric(t, resp.Metrics)
}

func TestPostInstrumentation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	}))
	defer ts.Close()

	resp, err := Post(ts.URL, "", nil)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	validateResponseMetric(t, resp.Metrics)
}
