import ctypes


class DynamicArray:
    """DynamicArray implements a simplified version of lists to demonstrate array resizing

    DynamicArray tracks its primitive operation "accounting credits" to illustrate the
    accounting method of analyzing amortized cost of operations on a data structure

    Underlying static storage still uses a ctypes array so that we can demonstrate resizing.

    Some obvious possible improvements to this implementation:
        1. Ensure we don't have a case where a 0-capacity underlying array gets "doubled"
            as part of an append operation, where 2 * capacity gets us another 0-capacity
            array that can't receive the element we want to append
        2. Set a minimum array size, so that a small array doesn't have to go through many
            costly resize operations too early in its lifecycle. For some reasonably small
            minimum array size, this is no-brainer memory/performance tradeoff to make.
    """

    DEFAULT_CAPACITY = 10
    _RESIZE_FACTOR = 2
    _SHRINK_THRESHOLD = 0.25

    def __init__(self):
        self._size = 0
        self.capacity = self.DEFAULT_CAPACITY
        self._arr = self._make_array(self.DEFAULT_CAPACITY)

    def __str__(self) -> str:
        arr_str = "["
        for i in range(self._size):
            arr_str += f"'{self._arr[i]}',"
        arr_str += "]"
        return f"DynamicArray: size={len(self)}, capacity={self.capacity}, {arr_str}"

    def __len__(self) -> int:
        return self._size

    def append(self, value: int) -> "DynamicArray":
        if self._size >= self.capacity:
            self._resize(self._size * self._RESIZE_FACTOR)

        self._arr[self._size] = value
        self._size += 1
        return self

    def _resize(self, new_capacity: int) -> None:
        new_arr = self._make_array(new_capacity)
        for i in range(self._size):
            new_arr[i] = self._arr[i]

        self._arr = new_arr
        self.capacity = new_capacity

    @staticmethod
    def _make_array(capacity: int) -> ctypes.Array:
        # https://docs.python.org/3/library/ctypes.html#arrays
        return (capacity * ctypes.c_int)()
