package algorithmanalysis

import "fmt"

const defaultCapacity = 10
const resizeFactor = 2
const shrinkThreshold float64 = 0.25

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
		Capacity:         10,
		Arr:              make([]int, 0, 10),
		OperationCredits: 0,
	}
}

func (da *DynamicArray) String() string {
	return fmt.Sprintf(
		"DynamicArray: Size=%d, Capacity=%d, OperationCredits=%d; underlying slice: len=%d cap=%d, location=%p %v,\n",
		da.Size, da.Capacity, da.OperationCredits, len(da.Arr), cap(da.Arr), da.Arr, da.Arr)
}

func (da *DynamicArray) Get(index int) int {
	return da.Arr[index]
}

func (da *DynamicArray) Pop() (*DynamicArray, int) {
	da.OperationCredits++

	sizeAfterPop := da.Size - 1
	if float64(sizeAfterPop)/float64(da.Capacity) < shrinkThreshold {
		newCapacity := da.Size * resizeFactor
		da = da.resize(da.Size, newCapacity)
	}

	value := da.Arr[len(da.Arr)-1]
	da.Arr = da.Arr[:len(da.Arr)-1]

	da.Size--
	// da.OperationCredits--
	da.OperationCredits--
	return da, value
}

func (da *DynamicArray) Append(value int) *DynamicArray {
	da.OperationCredits += 3

	// sizeAfterAppend := da.Size + 1
	if da.Size >= da.Capacity {
		newCapacity := da.Size * resizeFactor
		da = da.resize(da.Size, newCapacity)
	}

	da.Arr = append(da.Arr, value)
	da.OperationCredits--
	da.Size++

	return da
}

func (da *DynamicArray) resize(newSize, newCapacity int) *DynamicArray {
	fmt.Printf("RESIZE to capacity %d\n", newCapacity)
	newArr := make([]int, newSize, newCapacity)

	for i, v := range da.Arr {
		da.OperationCredits--
		newArr[i] = v
	}

	da.Arr = newArr
	da.Capacity = newCapacity

	return da
}
