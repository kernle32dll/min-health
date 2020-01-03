package health

import (
	"fmt"
	"net/http"
)

// Config contains the parsed config for the health check
type Config struct {
	Method string
	URL    string
}

// MakeRequest executes an health check, and returns true if it succeeded (or false if otherwise)
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
