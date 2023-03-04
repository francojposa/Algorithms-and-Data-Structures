package basicdatastructures

const DefaultQueueCapacity = 10

// ArrayQueue demonstrates an Array-based Queue implementation
//
// Underlying "static" storage still uses a slice, as Golang considers an array's size
// to be part of its type, so we cannot define the ArrayQueue with user-defined capacity
// when using an array as the underlying storage.
type ArrayQueue struct {
	data     []any
	capacity int
}

type ArrayQueueOpt func(queue *ArrayQueue)

func WithQueueCapacity(capacity int) ArrayQueueOpt {
	return func(queue *ArrayQueue) {
		queue.capacity = capacity
		queue.data = make([]any, 0, capacity)
	}
}

func NewArrayQueue(opts ...ArrayQueueOpt) *ArrayQueue {
	queue := &ArrayQueue{
		data:     make([]any, 0, DefaultQueueCapacity),
		capacity: DefaultQueueCapacity,
	}

	for _, opt := range opts {
		opt(queue)
	}
	return queue
}

func (aq *ArrayQueue) EnQueue(v any) error {
	if len(aq.data) >= aq.capacity {
		return QueueFullError{}
	}
	aq.data = append(aq.data, v)
	return nil
}

func (aq *ArrayQueue) DeQueue() (any, error) {
	if len(aq.data) == 0 {
		return nil, QueueEmptyError{}
	}
	v := aq.data[0]
	aq.data = aq.data[1:]
	return v, nil
}

func (aq *ArrayQueue) Len() int {
	return len(aq.data)
}
