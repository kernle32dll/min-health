package health

import (
	"fmt"
	"net/http"
)

// Config contains the parsed config for the health check
type Config struct {
	Method string
	URL    string
	Client *http.Client
}

// DoRequest executes an health check, and returns true if it succeeded (or false if otherwise)
func DoRequest(config *Config) bool {
	if config == nil {
		fmt.Println("no config provided")
		return false
	}

	req, err := http.NewRequest(config.Method, config.URL, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("error constructing request: %s", err))
		return false
	}

	client := config.Client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
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
