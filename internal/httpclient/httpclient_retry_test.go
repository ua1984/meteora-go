package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type HTTPClientRetryTestSuite struct {
	suite.Suite
}

func TestHTTPClientRetry(t *testing.T) {
	suite.Run(t, new(HTTPClientRetryTestSuite))
}

func (s *HTTPClientRetryTestSuite) TestExponentialBackoffOnRateLimiting() {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount <= 3 {
			// Simulate rate limiting for first 3 attempts
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rate limit exceeded"))
			return
		}
		// Success on 4th attempt
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client := New(server.URL, nil)
	client.maxRetries = 5 // Allow enough retries
	client.baseDelay = 10 * time.Millisecond // Short delay for testing
	client.maxDelay = 50 * time.Millisecond

	var result map[string]bool
	startTime := time.Now()
	err := client.Get(context.TODO(), "/test", nil, &result)
	duration := time.Since(startTime)

	s.NoError(err)
	s.Equal(map[string]bool{"success": true}, result)
	s.Equal(4, attemptCount) // 1 initial + 3 retries
	// Verify that some delay occurred (exponential backoff)
	s.True(duration >= 10*time.Millisecond, "Expected some delay due to backoff")
}

func (s *HTTPClientRetryTestSuite) TestExponentialBackoffOnServerErrors() {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount <= 2 {
			// Simulate server errors for first 2 attempts
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		// Success on 3rd attempt
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client := New(server.URL, nil)
	client.maxRetries = 3
	client.baseDelay = 5 * time.Millisecond
	client.maxDelay = 20 * time.Millisecond

	var result map[string]string
	err := client.Get(context.TODO(), "/test", nil, &result)

	s.NoError(err)
	s.Equal(map[string]string{"status": "ok"}, result)
	s.Equal(3, attemptCount) // 1 initial + 2 retries
}

func (s *HTTPClientRetryTestSuite) TestNoRetryOnClientErrors() {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		// Always return 404 (client error - should not retry)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}))
	defer server.Close()

	client := New(server.URL, nil)
	client.maxRetries = 3

	var result map[string]string
	err := client.Get(context.TODO(), "/test", nil, &result)

	s.Error(err)
	s.IsType(&APIError{}, err)
	s.Equal(1, attemptCount) // Should not retry on 404
}

func (s *HTTPClientRetryTestSuite) TestMaxRetriesExceeded() {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		// Always return 500
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	client := New(server.URL, nil)
	client.maxRetries = 2 // Only 2 retries
	client.baseDelay = 1 * time.Millisecond

	var result map[string]string
	err := client.Get(context.TODO(), "/test", nil, &result)

	s.Error(err)
	s.IsType(&APIError{}, err)
	s.Equal(3, attemptCount) // 1 initial + 2 retries
}

func (s *HTTPClientRetryTestSuite) TestCustomRetryConfiguration() {
	client := NewWithRetryConfig("http://example.com", nil, 10, 20*time.Millisecond, 100*time.Millisecond)

	s.Equal(10, client.maxRetries)
	s.Equal(20*time.Millisecond, client.baseDelay)
	s.Equal(100*time.Millisecond, client.maxDelay)
}

func (s *HTTPClientRetryTestSuite) TestContextCancellationDuringRetry() {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		// Always return 500 to trigger retries
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	client := New(server.URL, nil)
	client.maxRetries = 5
	client.baseDelay = 100 * time.Millisecond // Long delay to test cancellation
	client.maxDelay = 200 * time.Millisecond

	// Create a context that cancels quickly
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var result map[string]string
	err := client.Get(ctx, "/test", nil, &result)

	s.Error(err)
	s.Contains(err.Error(), "context canceled")
	// Should have made at least one attempt but less than all retries
	s.True(attemptCount >= 1 && attemptCount < 6, "Should have been canceled during retry")
}

func (s *HTTPClientRetryTestSuite) TestDefaultRetryConfiguration() {
	client := New("http://example.com", nil)

	s.Equal(DefaultMaxRetries, client.maxRetries)
	s.Equal(DefaultBaseDelay, client.baseDelay)
	s.Equal(DefaultMaxDelay, client.maxDelay)
}