package algorithmanalysis

// MaxSumSubArraySlow is a naive implementation of the Max Sum Subarray Problem
// running in O(n^3) time by enumerating all n^2 possible subarrays and summing
// the m elements of each of those subarrays
func MaxSumSubArraySlow(a []int) []int {
	maxStart, maxEnd := 0, 0
	maxSum := 0 // base case subarray is empty with sum 0

	for end := 1; end < len(a); end++ {
		for start := 0; start < end; start++ {
			sum := 0
			for _, val := range a[start:end] {
				sum += val
			}
			if sum > maxSum {
				maxStart, maxEnd, maxSum = start, end, sum
			}
		}
	}

	return a[maxStart:maxEnd]
}

// MaxSumSubArrayFaster is a slightly-less naive implementation of the Max Sum Subarray Problem
// running in O(n^2) time by iterating once to build an array s of the sums of a[0..i] for i < n,
// then enumerating all n^2 possible subarrays, but calculating the sums in constant time by
// using the array of sums: sum(a[i..j]) = s[j] - s[i]
func MaxSumSubArrayFaster(a []int) []int {
	accumulatedSums := make([]int, len(a))

	// iterate once to build an array of "accumulated" sums
	// where element i is equal to sum of a[0] through a[i] inclusive;
	// * contributes O(n) runtime
	// * eliminates duplicated summation iterations from occurring inside an O(n^2)
	//  outer loop, which causes the slower version to be O(n^3)
	for i := 0; i < len(a); i++ {
		if i == 0 {
			accumulatedSums[i] = a[i]
		} else {
			accumulatedSums[i] = accumulatedSums[i-1] + a[i]
		}

	}

	maxAccumStart, maxAccumEnd := 0, 0
	maxSum := 0 // base case subarray is empty with sum 0

	// * enumerate all contiguous subarrays; contributes O(n^2) runtime
	// * each contiguous subarray's sum is calculated in constant time using accumulatedSums
	for end := 1; end < len(accumulatedSums); end++ {
		for start := 0; start < end; start++ {
			// sum of subarray is equal to:
			// * accumulated sum of a from a[0] to a[end] inclusive, minus
			// * accumulated sum of a from a[0] to a[start], inclusive
			sum := accumulatedSums[end] - accumulatedSums[start]
			if sum > maxSum {
				maxAccumStart, maxAccumEnd, maxSum = start, end, sum
			}
		}
	}
	// indices based on array of sums need bumped up
	// Ex: maxSum occurred when maxAccumStart = 1 and maxAccumEnd = 3
	// accumulatedSums[1] is sum of a from a[0] to a[1] inclusive - 2 elements
	// accumulatedSums[3] is sum of a from a[0] to a[3] inclusive - 4 elements
	//
	// So accumulatedSums[3] - accumulatedSums[1] is equal to:
	// * sum of a from a[0] to a[3] inclusive, minus
	// * sum of a from a[0] to a[1] inclusive,
	// which is equivalent to the sum of a from a[2] to a[3] inclusive.
	// We want to return the slice of a which includes elements 2 and 3; which is a[2:4]
	// so from maxAccumStart = 1, maxAccumEnd = 3, we needed indices 2 and 4 to slice a
	return a[maxAccumStart+1 : maxAccumEnd+1]
}

// MaxSumSubArrayFastest is an optimized implementation of the Max Sum Subarray Problem
// running in O(n) time by iterating once to build an array m of the maximum subarray sum
// for subarrays of a[0..i] for i < n where the subarray includes the element at index i,
// then returning the maximum of those maximums.
//
// This is enabled by the following observation:
// If the sum of the max sum subarray ending at and including element at index i is not greater than 0,
// then that subarray's sum can trivially be improved upon by replacing it with the empty set, which has sum 0.
// That is, if the max sum subarray ending at and including element at index i does not contribute a positive
// value to our max sum, then it is best to leave the subarray out completely.
// To illustrate, use the array a = [-2, -4, 3, -1, 5, 6, -7, -2, 4, -3, 2] Base case empty set sum is 0.
// 1. subarray space is a[0..0] = [-2]
//   - The preceding subarray sum 0, plus this element -2, has sum -2
//   - This can trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0]
//
// 2. subarray space is a[0..1] = [-2, -4]
//   - The preceding subarray sum 0, plus this element -4, has sum -4
//   - This can trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0, 0]
//
// 3. subarray space is a[0..2] = [-2, -4, 3]
//   - The preceding subarray sum 0, plus this element 3, has sum 3
//   - This cannot trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0, 0, 3]
//
// 4. subarray space is a[0..3] = [-2, -4, 3, -1]
//   - The preceding subarray sum 3, plus this element -1, has sum 2
//   - This cannot trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0, 0, 3, 2]
//
// 5. subarray space is a[0..4] = [-2, -4, 3, -1, 5]
//   - The preceding subarray sum 2, plus this element 5, has sum 7
//   - This cannot trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0, 0, 3, 2, 7]
//
// 6. subarray space is a[0..5] = [-2, -4, 3, -1, 5, 6]
//   - The preceding subarray sum 7, plus this element 6, has sum 13
//   - This cannot trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0, 0, 3, 2, 7, 13]
//
// 7. subarray space is a[0..6] = [-2, -4, 3, -1, 5, 6, -7]
//   - The preceding subarray sum 13, plus this element -7, has sum 6
//   - This cannot trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0, 0, 3, 2, 7, 13, 6]
//
// 8. subarray space is a[0..7] = [-2, -4, 3, -1, 5, 6, -7, -2]
//   - The preceding subarray sum 6, plus this element -2, has sum 4
//   - This cannot trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0, 0, 3, 2, 7, 13, 6, 4]
//
// 9. subarray space is a[0..8] = [-2, -4, 3, -1, 5, 6, -7, -2, 4]
//   - The preceding subarray sum 4, plus this element 4, has sum 8
//   - This cannot trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0, 0, 3, 2, 7, 13, 6, 4, 8]
//
// 10. subarray space is a[0..9] = [-2, -4, 3, -1, 5, 6, -7, -2, 4, -3]
//   - The preceding subarray sum 8, plus this element -3, has sum 5
//   - This cannot trivially be improved upon by starting over with the empty set of sum 0
//   - Array of max sums of subarray space is [0, 0, 3, 2, 7, 13, 6, 4, 8, 5]
//
// The maximum possible sum of all subarrays of a is the max of our array of local subarray,
// maximums m =[0, 0, 3, 2, 7, 13, 6, 4, 8, 5], which is 13.
func MaxSumSubArrayFastest(a []int) []int {
	maxSumsEndingWithI := make([]int, len(a))

	emptySubArraySum := 0 // base case subarray is empty with sum 0
	maxSum := emptySubArraySum
	currentSubArrayStart := 0
	maxStart, maxEnd := 0, 0

	for i := 0; i < len(a); i++ {

		// look up max possible sum of the preceding subarray ending at i-1, inclusive
		var maxSumEndingBeforeI int
		if i == 0 {
			maxSumEndingBeforeI = 0
		} else {
			maxSumEndingBeforeI = maxSumsEndingWithI[i-1]
		}

		currentSubArrayMaxSum := maxSumEndingBeforeI + a[i]
		if currentSubArrayMaxSum <= emptySubArraySum {
			// If max possible sum of the subarray ending at i-1, inclusive, plus a[i]
			// is less than or equal to than the base case empty subarray sum 0.
			// then it's as good (== 0) or better (<0) to start over with an empty subarray
			// than to extend the previous max sum subarray to include a[i].
			// So the max sum of a subarray ending with a[i] is the empty subarray, with sum 0
			maxSumsEndingWithI[i] = emptySubArraySum
			// The next possible candidate subarray will not include a[i];
			// the earliest starting point would be at a[i+1].
			currentSubArrayStart = i + 1
		} else {
			// Otherwise, the subarray ending at i-1, inclusive, now extended to include a[i],
			// is still a candidate for max sum subarray.
			// Even if the sum of the subarray ending with a[i] is less than the one ending with a[i-1],
			// as long as it's still positive, the sum could still be improved with an upcoming value
			maxSumsEndingWithI[i] = currentSubArrayMaxSum
			// The candidate subarray is an extension of the previous; currentSubArrayStart stays.
		}
		if currentSubArrayMaxSum > maxSum {
			// record bounds of the best known max sum subarray so far
			maxSum = currentSubArrayMaxSum
			maxStart, maxEnd = currentSubArrayStart, i
		}
	}

	return a[maxStart : maxEnd+1]
}
