package httpclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want   string
	}{
		{
			name:   "should format status 404 with textual body",
			status: 404,
			body:   "Not Found",
			want:   "meteora API error: status 404: Not Found",
		},
		{
			name:   "should format status 500 with empty body",
			status: 500,
			body:   "",
			want:   "meteora API error: status 500: ",
		},
		{
			name:   "should format status 400 with JSON body",
			status: 400,
			body:   `{"error":"bad request"}`,
			want:   `meteora API error: status 400: {"error":"bad request"}`,
		},
		{
			name:   "should extract message field from API error JSON",
			status: 400,
			body:   `{"message":"invalid pool address"}`,
			want:   `meteora API error: status 400: invalid pool address`,
		},
		{
			name:   "should extract message field from API error JSON with extra fields",
			status: 422,
			body:   `{"message":"page must be >= 1","code":422}`,
			want:   `meteora API error: status 422: page must be >= 1`,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			err := &APIError{
				StatusCode: tt.status,
				Body:       tt.body,
			}

			// Act
			got := err.Error()

			// Assert
			assert.Equal(t, tt.want, got)
		})
	}
}
