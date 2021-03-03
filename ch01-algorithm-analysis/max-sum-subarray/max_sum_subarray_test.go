package algorithmanalysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumSubarray(t *testing.T) {
	arr1 := []int{-2, -4, 3, -1, 5, 6, -7, -2, 4, -3, 2}
	arr2 := []int{2, -3, 4, -1, -2, 1, 5, -3}
	max1 := MaxSumSubArraySlow(arr1)
	max2 := MaxSumSubArraySlow(arr2)
	assert.Equal(t, 13, max1)
	assert.Equal(t, 7, max2)

	max1 = MaxSumSubArrayFaster(arr1)
	max2 = MaxSumSubArrayFaster(arr2)
	assert.Equal(t, 13, max1)
	assert.Equal(t, 7, max2)

	max1 = MaxSumSubArrayFastest(arr1)
	max2 = MaxSumSubArrayFastest(arr2)
	assert.Equal(t, 13, max1)
	assert.Equal(t, 7, max2)
}
