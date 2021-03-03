from dynamic_array import DynamicArray


def test_dynamic_array():
    da = DynamicArray()

    # Fill dynamic array up to initial capacity
    for i in range(10):
        da.append(i)
        assert da[i] == i
        assert da._operation_credits > 0

    # Pop pop pop til it's gone
    for i in range(9, -1, -1):
        _, val = da.pop()
        assert val == i
        assert da._operation_credits > 0

    # Fill dynamic array back up to initial capacity
    for i in range(10):
        da.append(i)
        assert da[i] == i
        assert da._operation_credits > 0

    # Fill dynamic array up to 2x initial capacity
    for i in range(10, 20):
        da.append(i)
        assert da[i] == i
        assert da._operation_credits > 0
