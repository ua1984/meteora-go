package meteora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "should return pointer to non-empty string",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "should return pointer to empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			input := tt.input

			// Act
			result := String(input)

			// Assert
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "should return pointer to positive integer",
			input:    42,
			expected: 42,
		},
		{
			name:     "should return pointer to negative integer",
			input:    -10,
			expected: -10,
		},
		{
			name:     "should return pointer to zero",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			input := tt.input

			// Act
			result := Int(input)

			// Assert
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    bool
		expected bool
	}{
		{
			name:     "should return pointer to true",
			input:    true,
			expected: true,
		},
		{
			name:     "should return pointer to false",
			input:    false,
			expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			input := tt.input

			// Act
			result := Bool(input)

			// Assert
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}