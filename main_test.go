package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestParseArguments(t *testing.T) {
	defaultConfig := &Config{
		Method: http.MethodGet,
		URL:    "http://localhost/",
	}

	tests := []struct {
		name string
		args []string
		want *Config
	}{
		{"empty", nil, defaultConfig},
		{"missing-value", []string{"--missing-value"}, defaultConfig},
		{"not-a-param", []string{"not-a-param"}, defaultConfig},
		{"invalid-url", []string{"--url=" + string(0x7f)}, defaultConfig},
		{"valid-method", []string{"--method=POST"}, &Config{
			Method: http.MethodPost,
			URL:    defaultConfig.URL,
		}},
		{"valid-url", []string{"--url=http://localhost:5000/health"}, &Config{
			Method: defaultConfig.Method,
			URL:    "http://localhost:5000/health",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseArguments(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArguments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeRequest(t *testing.T) {
	defaultConfig := &Config{
		Method: http.MethodGet,
		URL:    "http://localhost/",
	}

	tests := []struct {
		name   string
		config *Config
		server *httptest.Server
		want   bool
	}{
		{"200", defaultConfig, statusCodeAnsweringServer(http.StatusOK), true},
		{"404", defaultConfig, statusCodeAnsweringServer(http.StatusNotFound), false},
		{"no-route", &Config{
			Method: defaultConfig.Method,
			URL:    "",
		}, nil, false},
		{"invalid-url", &Config{
			Method: defaultConfig.Method,
			URL:    string(0x7f),
		}, nil, false},
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

			if got := MakeRequest(tt.config); got != tt.want {
				t.Errorf("MakeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func statusCodeAnsweringServer(statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(statusCode)
	}))
}
