package binarysearchtrees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	a        []int
	k        int
	expected int
}{
	{
		a:        []int{},
		k:        1,
		expected: BinarySearchNotfound,
	},
	{
		a:        []int{0},
		k:        1,
		expected: BinarySearchNotfound,
	},
	{
		a:        []int{1},
		k:        1,
		expected: 0,
	},
	{
		a:        []int{2, 4, 7, 8, 8, 9, 13, 14, 14, 20, 25},
		k:        1,
		expected: BinarySearchNotfound,
	},
	{
		a:        []int{2, 4, 7, 8, 8, 9, 13, 14, 14, 20, 25},
		k:        9,
		expected: 5,
	},
	{
		a:        []int{2, 4, 7, 8, 8, 9, 13, 14, 14, 20, 25},
		k:        7,
		expected: 2,
	},
	{
		a:        []int{2, 4, 7, 8, 8, 9, 13, 14, 14, 20, 25},
		k:        20,
		expected: 9,
	},
	{
		a:        []int{2, 4, 7, 8, 8, 9, 13, 14, 14, 20, 25},
		k:        2,
		expected: 0,
	},
	{
		a:        []int{2, 4, 7, 8, 8, 9, 13, 14, 14, 20, 25},
		k:        25,
		expected: 10,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, BinarySearch(test.a, test.k))
	}
}
