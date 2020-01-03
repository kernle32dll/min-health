package main

import (
	"github.com/kernle32dll/min-health/health"

	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	config := ParseArguments(os.Args[1:])

	if health.DoRequest(config) {
		os.Exit(0)
	}

	os.Exit(1)
}

// ParseArguments parses the given arguments for a health check configuration
func ParseArguments(args []string) *health.Config {
	result := &health.Config{
		Method: http.MethodGet,
		URL:    "http://localhost/",
	}

	for _, arg := range args {
		split := strings.Split(arg, "=")

		if !strings.HasPrefix(arg, "--") || len(split) != 2 {
			fmt.Println(fmt.Sprintf("ignoring argument %q: not a proper argument", arg))
			continue
		}

		key := strings.Trim(split[0], "--")
		param := split[1]

		switch strings.ToLower(key) {
		case "method":
			result.Method = param
		case "url":
			if _, err := url.Parse(param); err != nil {
				fmt.Println(fmt.Sprintf("ignoring argument %q: invalid url", arg))
				continue
			}

			result.URL = param
		}
	}

	return result
}
