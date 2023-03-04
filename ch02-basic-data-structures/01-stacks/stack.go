package basicdatastructures

type Stack interface {
	Push(v any) error
	Pop() (any, error)
	Len() int
}

type StackFullError struct{}

func (e StackFullError) Error() string {
	return "stack full"
}

type StackEmptyError struct{}

func (e StackEmptyError) Error() string {
	return "stack empty"
}
