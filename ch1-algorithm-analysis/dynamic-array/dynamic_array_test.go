package algorithmanalysis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {
	assert := assert.New(t)
	da := NewDynamicArray()

	t.Run("Fill dynamic array up to initial capacity", func(t *testing.T) {
		for i := 0; i <= defaultCapacity; i++ {
			fmt.Print(da)
			da.Append(i)
			fmt.Print(da)
			assertDynamicArrayState(t, assert, da, i, i, i+1, defaultCapacity)
		}
	})

	t.Run("Pop pop pop til it's gone", func(t *testing.T) {
		for i := len(da.Arr) - 1; i >= 0; i-- {
			fmt.Print(da)
			da.Pop()
			fmt.Print(da)
		}
	})

	t.Run("Fill dynamic array up to initial capacity", func(t *testing.T) {
		for i := 0; i < defaultCapacity; i++ {
			fmt.Print(da)
			da.Append(i)
			fmt.Print(da)
			assertDynamicArrayState(t, assert, da, i, i, i+1, defaultCapacity)
		}
	})

	t.Run("Trigger dynamic array expansion, fill up to 2x initial capacity", func(t *testing.T) {
		for i := defaultCapacity; i < 2*defaultCapacity; i++ {
			fmt.Print(da)
			da.Append(i)
			fmt.Print(da)
			assertDynamicArrayState(t, assert, da, i, i, i+1, 2*defaultCapacity)
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
	// assert.Equal(expectedElement, da.Get(index))
	// assert.Equal(expectedSize, da.Size)
	// assert.Equal(expectedCapacity, da.Capacity)
	assert.Greater(da.OperationCredits, 0)
}
