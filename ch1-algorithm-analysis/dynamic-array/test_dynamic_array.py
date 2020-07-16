from dynamic_array import DynamicArray


def test_dynamic_array():
    da = DynamicArray()

    # Fill dynamic array up to initial capacity
    for i in range(10):
        da.append(i)
        assert da._operation_credits > 0
