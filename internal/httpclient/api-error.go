package httpclient

import (
	"encoding/json"
	"fmt"
)

// APIError represents an error response from the Meteora API.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	if msg := parseAPIErrorMessage(e.Body); msg != "" {
		return fmt.Sprintf("meteora API error: status %d: %s", e.StatusCode, msg)
	}
	return fmt.Sprintf("meteora API error: status %d: %s", e.StatusCode, e.Body)
}

// parseAPIErrorMessage attempts to extract the "message" field from a JSON error body.
// Returns an empty string if the body is not a valid API error JSON.
func parseAPIErrorMessage(body string) string {
	var apiErr struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal([]byte(body), &apiErr); err == nil && apiErr.Message != "" {
		return apiErr.Message
	}
	return ""
}
