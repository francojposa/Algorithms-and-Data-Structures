package basicdatastructures

const DefaultCapacity = 10

// ArrayStack demonstrates an Array-based Stack implementation
//
// Underlying "static" storage still uses a slice, as Golang considers an array's size
// to be part of its type, so we cannot define the ArrayStack with user-defined capacity
// when using an array as the underlying storage.
type ArrayStack struct {
	capacity int
	data     []any
}

type ArrayStackOpt func(stack *ArrayStack)

func WithStackCapacity(capacity int) ArrayStackOpt {
	return func(stack *ArrayStack) {
		stack.capacity = capacity
		stack.data = make([]any, 0, capacity)
	}
}

func NewArrayStack(opts ...ArrayStackOpt) *ArrayStack {
	stack := &ArrayStack{
		capacity: DefaultCapacity,
		data:     make([]any, 0, DefaultCapacity),
	}
	for _, opt := range opts {
		opt(stack)
	}
	return stack
}

func (as *ArrayStack) Push(v any) error {
	if len(as.data) == as.capacity {
		return StackFullError{}
	}
	as.data = append(as.data, v)
	return nil
}

func (as *ArrayStack) Pop() (any, error) {
	if len(as.data) == 0 {
		return nil, StackEmptyError{}
	}
	v := as.data[len(as.data)-1]
	as.data = as.data[:len(as.data)-1]
	return v, nil
}

func (as *ArrayStack) Len() int {
	return len(as.data)
}
