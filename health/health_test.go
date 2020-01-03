package health_test

import (
	"github.com/kernle32dll/min-health/health"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	defaultConfig := &health.Config{
		Method: http.MethodGet,
		URL:    "http://localhost/",
	}

	tests := []struct {
		name   string
		config *health.Config
		server *httptest.Server
		want   bool
	}{
		{"200", defaultConfig, statusCodeAnsweringServer(http.StatusOK), true},
		{"404", defaultConfig, statusCodeAnsweringServer(http.StatusNotFound), false},
		{"no-route", &health.Config{
			Method: defaultConfig.Method,
			URL:    "",
		}, nil, false},
		{"invalid-url", &health.Config{
			Method: defaultConfig.Method,
			URL:    string(0x7f),
		}, nil, false},
		{"custom-http-client", &health.Config{
			Method: defaultConfig.Method,
			URL:    defaultConfig.URL,
			Client: &http.Client{},
		}, statusCodeAnsweringServer(http.StatusOK), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.server != nil {
				defer tt.server.Close()

				// copy config and change url
				config := *tt.config
				config.URL = "http://" + tt.server.Listener.Addr().String()
				tt.config = &config
			}

			if got := health.DoRequest(tt.config); got != tt.want {
				t.Errorf("DoRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func statusCodeAnsweringServer(statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(statusCode)
	}))
}
