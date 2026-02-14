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
