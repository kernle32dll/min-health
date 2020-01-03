package main

import (
	"github.com/kernle32dll/min-health/health"

	"net/http"
	"reflect"
	"testing"
)

func TestParseArguments(t *testing.T) {
	defaultConfig := &health.Config{
		Method: http.MethodGet,
		URL:    "http://localhost/",
	}

	tests := []struct {
		name string
		args []string
		want *health.Config
	}{
		{"empty", nil, defaultConfig},
		{"missing-value", []string{"--missing-value"}, defaultConfig},
		{"not-a-param", []string{"not-a-param"}, defaultConfig},
		{"invalid-url", []string{"--url=" + string(0x7f)}, defaultConfig},
		{"valid-method", []string{"--method=POST"}, &health.Config{
			Method: http.MethodPost,
			URL:    defaultConfig.URL,
		}},
		{"valid-url", []string{"--url=http://localhost:5000/health"}, &health.Config{
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
