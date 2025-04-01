package httptestutil

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// RequestBuilder is a struct that helps to build HTTP requests
type RequestBuilder struct {
	Method      string
	Path        string
	Query       map[string]string
	Body        io.Reader
	ContentType string
}

// NewRequest creates a new HTTP request based on the provided builder
func NewRequest(t *testing.T, builder RequestBuilder) *http.Request {
	t.Helper()

	req := httptest.NewRequest(builder.Method, builder.Path, builder.Body)
	req.Header.Set("Content-Type", builder.ContentType)
	if builder.Query != nil {
		v := req.URL.Query()
		for key, value := range builder.Query {
			v.Set(key, value)
		}
		req.URL.RawQuery = v.Encode()
	}

	return req
}
