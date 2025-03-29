package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStrobogrammatic(t *testing.T) {
	tests := []struct{
		input string
		expected bool
	}{
		{
			input: "213",
			expected: false,
		},
		{
			input: "1",
			expected: true,
		},
		{
			input: "6",
			expected: false,
		},
		{
			input: "9",
			expected: false,
		},
		{
			input: "8",
			expected: true,
		},
		{
			input: "0",
			expected: true,
		},
		{
			input: "69",
			expected: true,
		},
		{
			input: "27",
			expected: false,
		},
		{
			input: "16891",
			expected: true,
		},
	}

	for _, test := range tests {
		isStrobogrammatic := isStrobogrammatic(test.input)

		assert.Equal(t, test.expected, isStrobogrammatic)
	}
}
