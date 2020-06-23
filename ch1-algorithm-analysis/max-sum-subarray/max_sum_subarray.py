from typing import Iterable


def max_sum_subarray_slow(a: Iterable[int]) -> int:
    """
    max_sum_subarray_slow is a naive implementation of the Max Sum Subarray
    Problem, running in O(n^3) time by enumerating all n^2 possible
    subarrays and summing the m elements of each of those subarrays
    """
    # the base case subarray is empty with sum 0
    max = 0
    for start_index in range(len(a)):
        # we will enumerate all subarrays that start at index 0,
        # then all subarrays that start at index 1, etc...
        for end_index in range(start_index, len(a)):
            # when start index is 0, we make subarrays a[0..0] up to a[0..len(a) - 1]
            # when start index is 1, we make subarrays a[1..1] up to a[1..len(a) - 1]
            # and so on. There is no need to "look back" from a start index of 1
            # to cover the subarray a[1..0] because a[0..1] was already covered on the
            # previous outer loop iteration and order of elements does not matter
            sub_array_sum = sum(a[start_index : end_index + 1])
            # sum all elements a[start index...end index]
            if sub_array_sum > max:
                max = sub_array_sum
    return max


def max_sum_subarray_faster(a: Iterable[int]) -> int:
    """
    max_sum_subarray_faster is a slightly-less naive implementation of the Max Sum Subarray Problem
    running in O(n^2) time by iterating once to build an array s of the sums of a[0..i] for i < n,
    then enumerating all n^2 possible subarrays, but calculating the sums in constant time by
    using the array of sums: sum(a[i..j]) = s[j] - s[i]
    """
    # build an array of the sums of a[0..i] for i < n
    sums = []
    cumulativeSum = 0
    for val in a:  # there's almost certainly a cooler a map/lambda way to do this
        cumulativeSum += val
        sums.append(cumulativeSum)

    max = 0  # the base case subarray is empty with sum 0

    for start_index in range(len(a)):
        # we will enumerate all subarrays that start at index 0,
        # then all subarrays that start at index 1, etc...
        for end_index in range(start_index + 1, len(a)):
            # when start index is 0, we make subarrays a[0..0] up to a[0..len(a) - 1]
            # when start index is 1, we make subarrays a[1..1] up to a[1..len(a) - 1]
            # and so on. There is no need to "look back" from a start index of 1
            # to cover the subarray a[1..0] because a[0..1] was already covered on the
            # previous outer loop iteration and order of elements does not matter for sums
            sum = sums[end_index] - sums[start_index]
            if sum > max:
                max = sum

    return max


def max_sum_subarray_fastest(a: Iterable[int]) -> int:
    """
    max_sum_subarray_fastest is an optimized implementation of the Max Sum Subarray Problem
    running in O(n) time by iterating once to build an array m of the maximum subarray sum
    for subarrays of a[0..i] for i < n where the subarray includes the element at index i,
    then returning the maximum of those maximums.

    This is enabled by the following observation:
    If the sum of the max sum subarray ending at and including element at index i is not greater than 0,
    then that subarray's sum can trivially be improved upon by replacing it with the empty set, which has sum 0.
    That is, if the max sum subarray ending at and including element at index i does not contribute a positive
    value to our max sum, then it is best to leave the subarray out completely.
    """
    BASE_CASE_EMPTY_SUBARRAY_SUM = 0
    overall_max_subarray_sum = BASE_CASE_EMPTY_SUBARRAY_SUM

    max_suffix_sums = []
    # build an array of the max suffix sums of all subarrays a[i..t] (inclusive) for 0 <= i <= t.
    for i, val in enumerate(a):
        previous_max_suffix_sum = (
            max_suffix_sums[i - 1] if i != 0 else BASE_CASE_EMPTY_SUBARRAY_SUM
        )

        # What do we get if we take the max sum subarray ending at i - 1 and add element i?
        # Is it possible to improve upon that preceding subarray?
        current_max_suffix_sum = previous_max_suffix_sum + val

        # If the max sum of a subarray that ends with the element at index i is negative, then
        # there is nothing to be gained by adding it onto the max sum subarray that ended at i - 1.
        # We are better off letting that preceding max sum subarray end and starting over with an empty subarray.
        # Therefore, the max sum of a subarray ending at index i is the empty subarray, with sum 0
        max_suffix_sums.append(
            max(current_max_suffix_sum, BASE_CASE_EMPTY_SUBARRAY_SUM)
        )

        # Keep track of the overall max subarray sum found so far
        overall_max_subarray_sum = max(overall_max_subarray_sum, current_max_suffix_sum)

    return overall_max_subarray_sum
