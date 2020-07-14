package algorithmanalysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {
	assert := assert.New(t)
	da := NewDynamicArray()

	t.Run("Fill dynamic array up to initial capacity", func(t *testing.T) {
		for i := 0; i < defaultCapacity; i++ {
			da.Append(i)
			// Assert that we have not run of out of our amortized operation credits
			assert.GreaterOrEqual(da.operationCredits, 0)
			// In reality, it doesn't matter if the credits are briefly below zero
			// as long as we are on average allocating enough credits to cover the cost
			// of many append/pop operations on a dynamic array.
			// The point is to illustrate that the averaged/amortized cost does not exceed
			// some constant factor per operation, meaning we have O(n) running time
		}
	})

	t.Run("Pop pop pop til it's gone", func(t *testing.T) {
		for i := da.Size - 1; i >= 0; i-- {
			da.Pop()
			// assert that we have not run of out of our amortized operation credits
			assert.GreaterOrEqual(da.operationCredits, 0)
		}
	})

	t.Run("Fill dynamic array back up to initial capacity", func(t *testing.T) {
		for i := 0; i < defaultCapacity; i++ {
			da.Append(i)
			// assert that we have not run of out of our amortized operation credits
			assert.GreaterOrEqual(da.operationCredits, 0)
		}
	})

	t.Run("Fill dynamic array up to 2x initial capacity", func(t *testing.T) {
		for i := defaultCapacity; i < 2*defaultCapacity; i++ {
			da.Append(i)
			// assert that we have not run of out of our amortized operation credits
			assert.GreaterOrEqual(da.operationCredits, 0)
		}
	})

}
