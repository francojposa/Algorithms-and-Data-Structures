package basicdatastructures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayQueue(t *testing.T) {

	t.Run("default capacity", func(t *testing.T) {

		inputs := [DefaultQueueCapacity + 2]int{0, 2, 4, 6, 8, 16, 32, 64, 128, 256, 512, 1024}
		queue := NewArrayQueue()

		for i, v := range inputs {
			err := queue.EnQueue(v)
			if i < DefaultQueueCapacity {
				assert.Nil(t, err)
				assert.Equal(t, i+1, queue.Len())
			} else {
				assert.ErrorIs(t, err, QueueFullError{})
				assert.Equal(t, DefaultQueueCapacity, queue.Len())
			}
		}

		reconstructedInputs := [DefaultQueueCapacity]any{}
		for i := 0; i < len(inputs); i++ {
			val, err := queue.DeQueue()
			if i < DefaultQueueCapacity {
				assert.Nil(t, err)
				reconstructedInputs[i] = val
			} else {
				assert.ErrorIs(t, err, QueueEmptyError{})
			}
		}
		for i := 0; i < DefaultQueueCapacity; i++ {
			assert.Equal(t, inputs[i], reconstructedInputs[i])
		}
	})

	t.Run("user-defined capacity", func(t *testing.T) {

		const capacity = 4
		inputs := [capacity + 2]int{0, 2, 4, 6, 8, 16}
		queue := NewArrayQueue(WithQueueCapacity(capacity))

		for i, v := range inputs {
			err := queue.EnQueue(v)
			if i < capacity {
				assert.Nil(t, err)
				assert.Equal(t, i+1, queue.Len())
			} else {
				assert.ErrorIs(t, err, QueueFullError{})
				assert.Equal(t, capacity, queue.Len())
			}
		}

		reconstructedInputs := [capacity]any{}
		for i := 0; i < len(inputs); i++ {
			val, err := queue.DeQueue()
			if i < capacity {
				assert.Nil(t, err)
				reconstructedInputs[i] = val
			} else {
				assert.ErrorIs(t, err, QueueEmptyError{})
			}
		}
		for i := 0; i < capacity; i++ {
			assert.Equal(t, inputs[i], reconstructedInputs[i])
		}
	})
}
