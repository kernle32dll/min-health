package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Config contains the parsed config for the health check
type Config struct {
	Method string
	URL    string
}

func main() {
	config := ParseArguments(os.Args[1:])

	if MakeRequest(config) {
		os.Exit(0)
	}

	os.Exit(1)
}

func MakeRequest(config *Config) bool {
	req, err := http.NewRequest(config.Method, config.URL, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("error constructing request: %s", err))
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Sprintf("error conducting request: %s", err))
		return false
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true
	}

	fmt.Println(fmt.Sprintf("unhealthy status code: %d", resp.StatusCode))
	return false
}

func ParseArguments(args []string) *Config {
	result := &Config{
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
