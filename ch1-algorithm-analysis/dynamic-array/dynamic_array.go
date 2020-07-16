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
// Go builtin slice functionalities: indexing with brackets, iterating with range, etc
//
// Some obvious possible improvements to this implementation:
//	1. Ensure we don't have a case where a 0-capacity underlying array gets "doubled"
//		as part of an append operation, where 2 * capacity gets us another 0-capacity
//		array that can't receive the element we want to append
//	2. Set a minimum array size, so that a small array doesn't have to go through many
//		costly resize operations too early in its lifecycle. For some reasonably small
//		minimum array size, this is no-brainer memory/performance tradeoff to make.
type DynamicArray struct {
	Size             int   // Number of actual elements
	Capacity         int   // Capacity of underlying static array
	Arr              []int // Static array storage
	operationCredits int   // Number of primitive operation credits built up
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		Size:             0,
		Capacity:         defaultCapacity,
		Arr:              make([]int, 0, defaultCapacity),
		operationCredits: 0,
	}
}

func (da *DynamicArray) String() string {
	return fmt.Sprintf(
		"DynamicArray: Size=%d, Capacity=%d, %v", da.Size, da.Capacity, da.Arr)
}

func (da *DynamicArray) Get(index int) int {
	return da.Arr[index]
}

// Append adds the given value to the end of a DynamicArray, doubling the allocated
// capacity of the underlying "static" storage array if the array is already full.
//
// At the beginning of the Append operation, we charge 3 Operation Credits:
//	* 1 credit will be consumed by the primitive operation of
//		writing the new value to the end of storage array
//	* 2 credits will be saved for a potential resizing of the array
//
// In the event of the array capacity expanding by a factor of 2, we must have already saved
// up enough OperationCredits to copy over contents of the existing array into the new array.
// Ex:
// 1. Initial State, we receive some array to be managed with dynamic storage
//		Array = [1, x]. Size 1, Capacity 2, OperationCredits 0
// 2. Append.
//		* 3 OperationCredits charged
//		* -1 OperationCredit spent appending the new element
//		* Array = [1, 2]. Size 2, Capacity 2, OperationCredits 2
// 3. Append, triggers resize.
//		* 3 OperationCredits charged
//		* -2 OperationCredit spent copying over the existing elements
//		* -1 OperationCredit spent appending the new element
//		* Array = [1, 2, 3, x]. Size 3, Capacity 4, OperationCredits 2
// 4. Append
//		* 3 OperationCredits charged
//		* -1 OperationCredit spent appending the new element
//		* Array = [1, 2, 3, 4]. Size 4, Capacity 4, OperationCredits 4
// 5. Append, triggers resize
//		* 3 OperationCredits charged
//		* -4 OperationCredit spent copying over the existing elements
//		* -1 OperationCredit spent appending the new element
//		* Array = [1, 2, 3, 4, 5, x, x, x]. Size 5, Capacity 8, OperationCredits 2
//... etc etc
func (da *DynamicArray) Append(value int) *DynamicArray {
	da.operationCredits += 3

	if da.Size >= da.Capacity {
		newCapacity := da.Size * resizeFactor
		da = da.resize(da.Size, newCapacity)
	}

	da.Arr = append(da.Arr, value)
	da.Size++
	da.operationCredits--
	return da
}

func (da *DynamicArray) Pop() (*DynamicArray, int) {
	da.operationCredits++

	sizeAfterPop := da.Size - 1
	if float64(sizeAfterPop)/float64(da.Capacity) < shrinkThreshold {
		newCapacity := da.Size * resizeFactor
		da = da.resize(da.Size, newCapacity)
	}

	value := da.Arr[len(da.Arr)-1]
	da.Arr = da.Arr[:len(da.Arr)-1]

	da.Size--
	da.operationCredits--
	return da, value
}

func (da *DynamicArray) resize(newSize, newCapacity int) *DynamicArray {
	newArr := make([]int, newSize, newCapacity)

	for i, v := range da.Arr {
		da.operationCredits--
		newArr[i] = v
	}

	da.Arr = newArr
	da.Capacity = newCapacity

	return da
}
