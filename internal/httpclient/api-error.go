package httpclient

import "fmt"

// APIError represents an error response from the Meteora API.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("meteora API error: status %d: %s", e.StatusCode, e.Body)
}
