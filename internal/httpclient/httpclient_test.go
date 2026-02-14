package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HTTPClientTestSuite struct {
	suite.Suite
}

func TestHTTPClient(t *testing.T) {
	suite.Run(t, new(HTTPClientTestSuite))
}

func (s *HTTPClientTestSuite) setupTestServer(method, wantURL string, status int, response any) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Equal(method, r.Method)
		s.Equal(wantURL, r.URL.String())
		s.Equal("application/json", r.Header.Get("Accept"))
		s.Equal("Meteora Go SDK/1.0.0", r.Header.Get("User-Agent"))

		w.WriteHeader(status)
		if str, ok := response.(string); ok {
			w.Write([]byte(str))
		} else {
			json.NewEncoder(w).Encode(response)
		}
	}))
}

func (s *HTTPClientTestSuite) TestGet() {
	tests := []struct {
		name       string
		path       string
		query      url.Values
		response   any
		status     int
		wantErr    bool
		errType    any
		wantURL    string
		wantResult any
	}{
		{
			name:       "should successfully perform GET request and decode JSON",
			path:       "test",
			response:   map[string]string{"foo": "bar"},
			status:     http.StatusOK,
			wantURL:    "/test",
			wantResult: map[string]string{"foo": "bar"},
		},
		{
			name:       "should successfully perform GET request with query parameters",
			path:       "test",
			query:      url.Values{"a": []string{"1"}, "b": []string{"2"}},
			response:   map[string]string{"foo": "bar"},
			status:     http.StatusOK,
			wantURL:    "/test?a=1&b=2",
			wantResult: map[string]string{"foo": "bar"},
		},
		{
			name:     "should return APIError for non-2xx status code",
			path:     "error",
			status:   http.StatusNotFound,
			response: "Not Found",
			wantErr:  true,
			errType:  &APIError{},
			wantURL:  "/error",
		},
		{
			name:     "should return error for invalid JSON response",
			path:     "bad-json",
			status:   http.StatusOK,
			response: "{invalid json}",
			wantErr:  true,
			wantURL:  "/bad-json",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			client := New(server.URL, nil)
			var result map[string]string

			// Act
			err := client.Get(context.TODO(), tt.path, tt.query, &result)

			// Assert
			if tt.wantErr {
				s.Error(err)
				if tt.errType != nil {
					s.IsType(tt.errType, err)
				}
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, result)
			}
		})
	}
}

func (s *HTTPClientTestSuite) TestPost() {
	tests := []struct {
		name       string
		path       string
		query      url.Values
		response   any
		status     int
		wantErr    bool
		errType    any
		wantURL    string
		wantResult any
	}{
		{
			name:       "should successfully perform POST request and decode JSON",
			path:       "test",
			response:   map[string]string{"foo": "bar"},
			status:     http.StatusOK,
			wantURL:    "/test",
			wantResult: map[string]string{"foo": "bar"},
		},
		{
			name:     "should return APIError for non-2xx status code",
			path:     "error",
			status:   http.StatusInternalServerError,
			response: "Internal Server Error",
			wantErr:  true,
			errType:  &APIError{},
			wantURL:  "/error",
		},
		{
			name:     "should return error for invalid JSON response",
			path:     "bad-json",
			status:   http.StatusOK,
			response: "{invalid json}",
			wantErr:  true,
			wantURL:  "/bad-json",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodPost, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			client := New(server.URL, nil)
			var result map[string]string

			// Act
			err := client.Post(context.TODO(), tt.path, tt.query, &result)

			// Assert
			if tt.wantErr {
				s.Error(err)
				if tt.errType != nil {
					s.IsType(tt.errType, err)
				}
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, result)
			}
		})
	}
}

func (s *HTTPClientTestSuite) TestNewClient() {
	s.Run("should use default client if nil is provided", func() {
		// Arrange
		want := "http://example.com"

		// Act
		client := New(want, nil)

		// Assert
		s.Equal(http.DefaultClient, client.httpClient)
		s.Equal(want, client.baseURL)
	})

	s.Run("should use provided client", func() {
		// Arrange
		custom := &http.Client{}
		want := "http://example.com"

		// Act
		client := New(want, custom)

		// Assert
		s.Equal(custom, client.httpClient)
		s.Equal(want, client.baseURL)
	})
}
