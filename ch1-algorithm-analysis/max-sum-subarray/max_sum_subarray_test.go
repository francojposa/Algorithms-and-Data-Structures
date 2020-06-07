package algorithmanalysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumSubarray(t *testing.T) {
	arr := []int{-2, -4, 3, -1, 5, 6, -7, -2, 4, -3, 2}
	max := MaxSumSubArraySlow(arr)
	assert.Equal(t, 13, max)

	max = MaxSumSubArrayFaster(arr)
	assert.Equal(t, 13, max)
}
