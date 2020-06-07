from max_sum_subarray import max_sum_subarray_slow, max_sum_subarray_faster


def test_max_sum_subarray():
    arr = [2, -4, 3, -1, 5, 6, -7, -2, 4, -3, 2]
    assert 13 == max_sum_subarray_slow(arr)
    assert 13 == max_sum_subarray_faster(arr)
