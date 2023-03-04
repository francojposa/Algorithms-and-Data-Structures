package basicdatastructures

type Queue interface {
	EnQueue(v any) error
	DeQueue() (any, error)
	Len() int
}

type QueueFullError struct{}

func (e QueueFullError) Error() string {
	return "queue full"
}

type QueueEmptyError struct{}

func (e QueueEmptyError) Error() string {
	return "queue empty"
}
