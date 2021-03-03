import ctypes
from typing import Tuple


class DynamicArray:
    """DynamicArray implements a simplified version of lists to demonstrate array resizing

    DynamicArray tracks its primitive operation "accounting credits" to illustrate the
    accounting method of analyzing amortized cost of operations on a data structure

    Underlying static storage uses a Python ctypes array so that we can demonstrate resizing.

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
        self._capacity = self.DEFAULT_CAPACITY
        self._arr = self._make_array(self.DEFAULT_CAPACITY)
        self._operation_credits = 0

    def __str__(self) -> str:
        arr_str = "["
        for i in range(self._size):
            arr_str += f"'{self._arr[i]}',"
        arr_str += "]"
        return f"DynamicArray: size={len(self)}, capacity={self._capacity}, {arr_str}"

    def __len__(self) -> int:
        return self._size

    def __getitem__(self, index: int):
        if not 0 <= index < self._size:
            raise IndexError("index out of bounds")
        return self._arr[index]

    def append(self, value: int) -> "DynamicArray":
        """
        append adds the given value to the end of a DynamicArray, doubling the allocated
        capacity of the underlying static storage array if the array is already full.

        At the beginning of the Append operation, we charge 3 Operation Credits:
        * 1 credit will be consumed by the primitive operation of
        writing the new value to the end of storage array
        * 2 credits will be saved for a potential resizing of the array

        In the event of the array capacity expanding by a factor of 2, we must have already saved
        up enough OperationCredits to copy over contents of the existing array into the new array.

        Ex:
        1. Initial State, we receive some array to be managed with dynamic storage
            * Array = [1, x]. Size 1, Capacity 2, OperationCredits 0
        2. Append.
            * 3 OperationCredit charged
            * -1 OperationCredit spent appending the new element
            * Array = [1, 2]. Size 2, Capacity 2, OperationCredits 2
        3. Append, triggers resize.
            * 3 OperationCredit charged
            * -2 OperationCredit spent copying over the existing elements
            * -1 OperationCredit spent appending the new element
            * Array = [1, 2, 3, x]. Size 3, Capacity 4, OperationCredits 2
        4. Append
            * 3 OperationCredit charged
            * -1 OperationCredit spent appending the new element
            * Array = [1, 2, 3, 4]. Size 4, Capacity 4, OperationCredits 4
        5. Append, triggers resize
            * 3 OperationCredit charged
            * -4 OperationCredit spent copying over the existing elements
            * -1 OperationCredit spent appending the new element
            * Array = [1, 2, 3, 4, 5, x, x, x]. Size 5, Capacity 8, OperationCredits 2
        etc etc
        """
        self._operation_credits += 3

        if self._size >= self._capacity:
            self._resize(self._size * self._RESIZE_FACTOR)

        self._arr[self._size] = value
        self._size += 1
        self._operation_credits -= 1
        return self

    def pop(self) -> Tuple["DynamicArray", int]:
        """"""
        self._operation_credits += 1

        size_after_pop = self._size - 1
        if size_after_pop / self._capacity < self._SHRINK_THRESHOLD:
            self._resize(self._size * self._RESIZE_FACTOR)

        val = self._arr[self._size - 1]

        self._size -= 1
        self._operation_credits -= 1
        return self, val

    def _resize(self, new_capacity: int) -> None:
        new_arr = self._make_array(new_capacity)
        for i in range(self._size):
            self._operation_credits -= 1
            new_arr[i] = self._arr[i]

        self._arr = new_arr
        self._capacity = new_capacity

    @staticmethod
    def _make_array(capacity: int) -> ctypes.Array:
        # https://docs.python.org/3/library/ctypes.html#arrays
        return (capacity * ctypes.c_int)()
