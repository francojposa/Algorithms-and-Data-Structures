package basicdatastructures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayStack(t *testing.T) {

	t.Run("default capacity", func(t *testing.T) {

		inputs := [DefaultStackCapacity + 2]int{0, 2, 4, 6, 8, 16, 32, 64, 128, 256, 512, 1024}
		stack := NewArrayStack()

		for i, v := range inputs {
			err := stack.Push(v)
			if i < DefaultStackCapacity {
				assert.Nil(t, err)
				assert.Equal(t, i+1, stack.Len())
			} else {
				assert.ErrorIs(t, err, StackFullError{})
				assert.Equal(t, DefaultStackCapacity, stack.Len())
			}
		}

		reconstructedInputs := [DefaultStackCapacity]any{}
		for i := 0; i < len(inputs); i++ {
			val, err := stack.Pop()
			if i < DefaultStackCapacity {
				assert.Nil(t, err)
				reconstructedInputs[DefaultStackCapacity-i-1] = val
			} else {
				assert.ErrorIs(t, err, StackEmptyError{})
			}
		}
		for i := 0; i < DefaultStackCapacity; i++ {
			assert.Equal(t, inputs[i], reconstructedInputs[i])
		}
	})

	t.Run("user-defined capacity", func(t *testing.T) {

		const capacity = 4
		inputs := [capacity + 2]int{0, 2, 4, 6, 8, 16}
		stack := NewArrayStack(WithStackCapacity(capacity))

		for i, v := range inputs {
			err := stack.Push(v)
			if i < capacity {
				assert.Nil(t, err)
				assert.Equal(t, i+1, stack.Len())
			} else {
				assert.ErrorIs(t, err, StackFullError{})
				assert.Equal(t, capacity, stack.Len())
			}
		}

		reconstructedInputs := [capacity]any{}
		for i := 0; i < len(inputs); i++ {
			val, err := stack.Pop()
			if i < capacity {
				assert.Nil(t, err)
				reconstructedInputs[capacity-i-1] = val
			} else {
				assert.ErrorIs(t, err, StackEmptyError{})
			}
		}
		for i := 0; i < capacity; i++ {
			assert.Equal(t, inputs[i], reconstructedInputs[i])
		}
	})
}
