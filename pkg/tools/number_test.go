package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHumanizeNumber(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "0",
			expected: "0",
		},
		{
			input:    "1",
			expected: "1",
		},
		{
			input:    "-1",
			expected: "-1",
		},
		{
			input:    "999",
			expected: "999",
		},
		{
			input:    "1001",
			expected: "1.001",
		},
		{
			input:    "-162000000",
			expected: "-162.000.000",
		},
	}

	for _, testCase := range testCases {
		actual := HumanizeNumber(testCase.input)
		assert.Equal(t, testCase.expected, actual)
	}
}
