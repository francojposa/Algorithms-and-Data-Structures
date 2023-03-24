package binarysearchtrees

const BinarySearchNotfound = -1

func BinarySearch(a []int, k int) int {
	return binarySearch(a, k, 0, len(a))
}

func binarySearch(a []int, k, low, high int) int {
	length := high - low

	if length == 0 {
		return BinarySearchNotfound
	}
	if length == 1 {
		if a[low] == k {
			return low
		} else {
			return BinarySearchNotfound
		}
	}

	midIdx := low + (length / 2)
	if a[midIdx] < k {
		return binarySearch(a, k, midIdx+1, high)
	} else if a[midIdx] > k {
		return binarySearch(a, k, low, midIdx)
	}
	return midIdx
}
