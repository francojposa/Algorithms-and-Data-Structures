package algorithmanalysis

// MaxSumSubArraySlow is a naive implementation of the Max Sum Subarray Problem
// running in O(n^3) time by enumerating all n^2 possible subarrays and summing
// the m elements of each of those subarrays
func MaxSumSubArraySlow(a []int) int {
	max := 0 // the base case subarray is empty with sum 0
	for startIndex := range a {
		// we will enumerate all subarrays that start at index 0,
		// then all subarrays that start at index 1, etc...
		for endIndex := startIndex + 1; endIndex < len(a); endIndex++ {
			// when start index is 0, we make subarrays a[0..0] up to a[0..len(a) - 1]
			// when start index is 1, we make subarrays a[1..1] up to a[1..len(a) - 1]
			// and so on. There is no need to "look back" from a start index of 1
			// to cover the subarray a[1..0] because a[0..1] was already covered on the
			// previous outer loop iteration and order of elements does not matter for sums
			sum := 0
			for i := startIndex; i < endIndex; i++ {
				// sum all elements a[start index...end index]
				sum += a[i]
				if sum > max {
					max = sum
				}
			}
		}
	}
	return max
}

// MaxSumSubArrayFaster is a slightly-less naive implementation of the Max Sum Subarray Problem
// running in O(n^2) time by iterating once to build an array s of the sums of a[0..i] for i < n,
// then enumerating all n^2 possible subarrays, but calculating the sums in constant time by
// using the array of sums: sum(a[i..j]) = s[j] - s[i]
func MaxSumSubArrayFaster(a []int) int {
	// build an array of the sums of a[0..i] for i < n
	sums := make([]int, len(a))
	cumulativeSum := 0
	for i, v := range a {
		cumulativeSum += v
		sums[i] = cumulativeSum
	}

	max := 0 // the base case subarray is empty with sum 0

	for startIndex := range a {
		// we will enumerate all subarrays that start at index 0,
		// then all subarrays that start at index 1, etc...
		for endIndex := startIndex + 1; endIndex < len(a); endIndex++ {
			// when start index is 0, we make subarrays a[0..0] up to a[0..len(a) - 1]
			// when start index is 1, we make subarrays a[1..1] up to a[1..len(a) - 1]
			// and so on. There is no need to "look back" from a start index of 1
			// to cover the subarray a[1..0] because a[0..1] was already covered on the
			// previous outer loop iteration and order of elements does not matter for sums
			sum := sums[endIndex] - sums[startIndex]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

// MaxSumSubArrayFastest is an optimized implementation of the Max Sum Subarray Problem
// running in O(n) time by iterating once to build an array m of the maximum subarray sum
// for subarrays of a[0..i] for i < n where the subarray includes the element at index i,
// then returning the maximum of those maximums.
//
// This is enabled by the following observation:
// if the sum of the max sum subarray ending at and including element at index i is not greater than 0,
// then that subarray's sum can trivially be improved upon by replacing it with the empty set, which has sum 0.
// That is, if the max sum subarray ending at and including element at index i does not contribute a positive
// value to our max sum, then it is best to leave the subarray out completely.
// To illustrate, use the array a = [-2, -4, 3, -1, 5, 6, -7, -2, 4, -3, 2] Base case maximum sum is 0.
// 1. subarray space is a[0..0] = [-2]
//		The preceding subarray sum 0, plus this element -2, has sum -2
//		This can trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0]
// 2. subarray space is a[0..1] = [-2, -4]
//		The preceding subarray sum 0, plus this element -4, has sum -4
//		This can trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0, 0]
// 3. subarray space is a[0..2] = [-2, -4, 3]
//		The preceding subarray sum 0, plus this element 3, has sum 3
//		This cannot trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0, 0, 3]
// 4. subarray space is a[0..3] = [-2, -4, 3, -1]
//		The preceding subarray sum 3, plus this element -1, has sum 2
//		This cannot trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0, 0, 3, 2]
// 5. subarray space is a[0..4] = [-2, -4, 3, -1, 5]
//		The preceding subarray sum 2, plus this element 5, has sum 7
//		This cannot trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0, 0, 3, 2, 7]
// 6. subarray space is a[0..5] = [-2, -4, 3, -1, 5, 6]
//		The preceding subarray sum 7, plus this element 6, has sum 13
//		This cannot trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0, 0, 3, 2, 7, 13]
// 7. subarray space is a[0..6] = [-2, -4, 3, -1, 5, 6, -7]
//		The preceding subarray sum 13, plus this element -7, has sum 6
//		This cannot trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0, 0, 3, 2, 7, 13, 6]
// 8. subarray space is a[0..7] = [-2, -4, 3, -1, 5, 6, -7, -2]
//		The preceding subarray sum 6, plus this element -2, has sum 4
//		This cannot trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0, 0, 3, 2, 7, 13, 6, 4]
// 9. subarray space is a[0..8] = [-2, -4, 3, -1, 5, 6, -7, -2, 4]
//		The preceding subarray sum 4, plus this element 4, has sum 8
//		This cannot trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0, 0, 3, 2, 7, 13, 6, 4, 8]
// 10. subarray space is a[0..9] = [-2, -4, 3, -1, 5, 6, -7, -2, 4, -3]
//		The preceding subarray sum 8, plus this element -3, has sum 5
//		This cannot trivially be improved upon by starting over with the empty set of sum 0
//		Array of max sums of subarray space is [0, 0, 3, 2, 7, 13, 6, 4, 8, 5]
// The maximum possible sum of all subarrays of a is the max of our array of local subarray,
// maximums m =[0, 0, 3, 2, 7, 13, 6, 4, 8, 5], which is 13.
func MaxSumSubArrayFastest(a []int) int {
	return 0
}
