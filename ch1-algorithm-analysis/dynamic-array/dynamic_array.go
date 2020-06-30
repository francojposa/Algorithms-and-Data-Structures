package algorithmanalysis

const defaultCapacity = 10

// DynamicArray implements a simplified version of slices to demonstrate array resizing
//
// Underlying "static" storage still uses a slice, as Golang considers an array's size
// to be part of its type, so we cannot define the DynamicArray with arr size = 10
// then resize the underlying array and return a DynamicArray with arr size = 20
type DynamicArray struct {
	Size     int   // Number of actual elements
	Capacity int   // Capacity of underlying static array
	arr      []int // Static array storage
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{Size: 0, Capacity: defaultCapacity, arr: make([]int, defaultCapacity)}
}
