package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client is a shared HTTP client used by all service clients.
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// New creates a new Client with the given base URL.
func New(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
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

// do performs a request and decodes the JSON response into result.
func (c *Client) do(req *http.Request, result any) error {
	// Clone the request to avoid modifying the original request.
	r := req.Clone(req.Context())
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "Meteora Go SDK/1.0.0")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Body:       string(body),
		}
	}

	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("decoding response: %w", err)
		}
	}

	return nil
}
