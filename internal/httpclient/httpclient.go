package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultMaxRetries = 5
	DefaultBaseDelay  = 100 * time.Millisecond
	DefaultMaxDelay   = 5 * time.Second
)

// Client is a shared HTTP client used by all service clients.
type Client struct {
	httpClient *http.Client
	baseURL    string
	maxRetries int
	baseDelay  time.Duration
	maxDelay   time.Duration
}

// New creates a new Client with the given base URL and default retry configuration.
func New(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		maxRetries: DefaultMaxRetries,
		baseDelay:  DefaultBaseDelay,
		maxDelay:   DefaultMaxDelay,
	}
}

// NewWithRetryConfig creates a new Client with custom retry configuration.
func NewWithRetryConfig(baseURL string, httpClient *http.Client, maxRetries int, baseDelay, maxDelay time.Duration) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Validate and set defaults for retry configuration
	if maxRetries <= 0 {
		maxRetries = DefaultMaxRetries
	}
	if baseDelay <= 0 {
		baseDelay = DefaultBaseDelay
	}
	if maxDelay <= 0 {
		maxDelay = DefaultMaxDelay
	}
	if maxDelay < baseDelay {
		maxDelay = baseDelay
	}

	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		maxRetries: maxRetries,
		baseDelay:  baseDelay,
		maxDelay:   maxDelay,
	}
}

// Get performs a GET request and decodes the JSON response into result.
func (c *Client) Get(ctx context.Context, path string, query url.Values, result any) error {
	u, err := url.JoinPath(c.baseURL, path)
	if err != nil {
		return fmt.Errorf("building URL: %w", err)
	}

	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	return c.do(req, result)
}

// Post performs a POST request and decodes the JSON response into result.
func (c *Client) Post(ctx context.Context, path string, query url.Values, result any) error {
	u, err := url.JoinPath(c.baseURL, path)
	if err != nil {
		return fmt.Errorf("building URL: %w", err)
	}

	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	return c.do(req, result)
}

// shouldRetry determines if a request should be retried based on the error or status code.
func (c *Client) shouldRetry(err error, statusCode int) bool {
	if err != nil {
		// Retry on network errors
		return true
	}

	// Retry on rate limiting (429) and server errors (5xx)
	if statusCode == http.StatusTooManyRequests || statusCode >= 500 {
		return true
	}

	return false
}

// calculateDelay calculates the exponential backoff delay with jitter.
func (c *Client) calculateDelay(attempt int) time.Duration {
	// Exponential backoff: baseDelay * 2^(attempt-1)
	delay := c.baseDelay * time.Duration(1<<(attempt-1))

	// Add jitter to avoid thundering herd problem
	jitter := time.Duration(rand.Int63n(int64(delay / 4)))
	delay += jitter

	// Cap at maxDelay
	if delay > c.maxDelay {
		return c.maxDelay
	}

	return delay
}

// do performs a request and decodes the JSON response into result.
func (c *Client) do(req *http.Request, result any) error {
	var lastErr error
	var lastBody []byte

	for attempt := 1; attempt <= c.maxRetries+1; attempt++ {
		// Check context at the start of each attempt
		select {
		case <-req.Context().Done():
			return fmt.Errorf("context canceled before attempt %d: %w", attempt, req.Context().Err())
		default:
		}

		// Clone the request to avoid modifying the original request.
		r := req.Clone(req.Context())
		r.Header.Set("Accept", "application/json")
		r.Header.Set("User-Agent", "Meteora Go SDK/1.0.0")
		resp, err := c.httpClient.Do(r)
		if err != nil {
			lastErr = fmt.Errorf("executing request (attempt %d): %w", attempt, err)
			if c.shouldRetry(err, 0) && attempt <= c.maxRetries {
				delay := c.calculateDelay(attempt)
				select {
				case <-req.Context().Done():
					return fmt.Errorf("context canceled during retry delay: %w", req.Context().Err())
				case <-time.After(delay):
					continue
				}
			}

			return lastErr
		}

		lastBody, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			lastErr = fmt.Errorf("reading response body: %w", err)
			return lastErr
		}

		statusCode := resp.StatusCode

		if statusCode < 200 || statusCode >= 300 {
			apiErr := &APIError{
				StatusCode: resp.StatusCode,
				Body:       string(lastBody),
			}
			lastErr = apiErr

			if c.shouldRetry(nil, statusCode) && attempt <= c.maxRetries {
				delay := c.calculateDelay(attempt)
				select {
				case <-req.Context().Done():
					return fmt.Errorf("context canceled during retry delay: %w", req.Context().Err())
				case <-time.After(delay):
					continue
				}
			}
			return lastErr
		}

		if result != nil {
			if err := json.Unmarshal(lastBody, result); err != nil {
				return fmt.Errorf("decoding response: %w", err)
			}
		}

		return nil
	}

	// If we exhausted all retries, return the last error
	if lastErr != nil {
		return lastErr
	}

	// This should not happen, but return a generic error just in case
	return fmt.Errorf("request failed after %d attempts", c.maxRetries+1)
}
