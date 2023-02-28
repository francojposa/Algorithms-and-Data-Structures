package algorithmanalysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input    []int
	expected []int
}{
	{
		input:    []int{-2, -4, 3, -1, 5, 6, -7, -2, 4, -3, 2},
		expected: []int{3, -1, 5, 6},
	},
	{
		input:    []int{2, -3, 4, -1, -2, 1, 5, -3},
		expected: []int{4, -1, -2, 1, 5},
	},
}

func TestMaxSumSubArraySlow(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, MaxSumSubArraySlow(test.input))
	}
}

func TestMaxSumSubArrayFaster(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, MaxSumSubArrayFaster(test.input))
	}
}

func TestMaxSumSubArrayFastest(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, MaxSumSubArrayFastest(test.input))
	}
}
