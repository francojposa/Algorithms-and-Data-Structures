package algorithmanalysis

const defaultCapacity = 10
const resizeFactor = 2

// DynamicArray implements a simplified version of slices to demonstrate array resizing
//
// DynamicArray tracks its primitive operation "accounting credits" to illustrate the
// accounting method of analyzing amortized cost of operations on a data structure
//
// Underlying "static" storage still uses a slice, as Golang considers an array's size
// to be part of its type, so we cannot define the DynamicArray with arr size = 10
// then resize the underlying array and return a DynamicArray with arr size = 20
// We pretend the slice is static for the purpose of demonstrating our own implementation
//
// Without generics in stable Go (yet), we also cannot directly expose the same interface a
// real slice offers, so consumers of DynamicArray need to access the Arr attribute for
// Go builtins slice functionality like indexing with brackets or iterating with range
type DynamicArray struct {
	Size             int   // Number of actual elements
	Capacity         int   // Capacity of underlying static array
	Arr              []int // Static array storage
	OperationCredits int   // Number of primitive operation credits built up
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		Size:             0,
		Capacity:         defaultCapacity,
		Arr:              make([]int, 0, defaultCapacity),
		OperationCredits: 0,
	}
}

func (da *DynamicArray) Get(index int) int {
	return da.Arr[index]
}

func (da *DynamicArray) Append(value int) *DynamicArray {
	da.OperationCredits += 3

	if da.Size == da.Capacity {
		da = da.resize(da.Capacity * resizeFactor)
	}

	da.Arr = append(da.Arr, value)
	da.OperationCredits--
	da.Size++

	return da
}

func (da *DynamicArray) resize(capacity int) *DynamicArray {
	newArr := make([]int, da.Size, capacity)

	for i, v := range da.Arr {
		da.OperationCredits--
		newArr[i] = v
	}

	da.Arr = newArr
	da.Capacity = capacity

	return da
}
