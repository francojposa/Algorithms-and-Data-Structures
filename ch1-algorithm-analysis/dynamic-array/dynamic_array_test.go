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
			assertDynamicArrayState(t, assert, da, i, i, i+1, defaultCapacity)
		}
	})

	t.Run("Trigger dynamic array expansion, fill up to 2x initial capacity", func(t *testing.T) {
		for i := defaultCapacity; i < 2*defaultCapacity; i++ {
			da.Append(i)
			assertDynamicArrayState(t, assert, da, i, i, i+1, 2*defaultCapacity)
		}
	})

	t.Run("Trigger dynamic array expansion, fill up to 2x previous capacity", func(t *testing.T) {
		for i := 2 * defaultCapacity; i < 4*defaultCapacity; i++ {
			da.Append(i)
			assertDynamicArrayState(t, assert, da, i, i, i+1, 4*defaultCapacity)
		}
	})
}

func assertDynamicArrayState(
	t *testing.T,
	assert *assert.Assertions,
	da *DynamicArray,
	index,
	expectedElement,
	expectedSize,
	expectedCapacity int,
) {
	t.Helper()
	assert.Equal(expectedElement, da.Get(index))
	assert.Equal(expectedSize, da.Size)
	assert.Equal(expectedCapacity, da.Capacity)
	assert.Greater(da.OperationCredits, 0)
}
