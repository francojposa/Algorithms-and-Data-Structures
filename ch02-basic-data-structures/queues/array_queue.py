from __future__ import annotations

from abc import ABC, abstractmethod
from typing import Any, Tuple


class Queue(ABC):
    @abstractmethod
    def __len__(self) -> int:
        pass

    @abstractmethod
    def is_empty(self) -> bool:
        """Return True if queue is empty; False otherwise"""

    @abstractmethod
    def enqueue(self, val: Any) -> Queue:
        """Insert value at end of queue"""

    @abstractmethod
    def dequeue(self) -> Tuple[Queue, Any]:
        """Remove value from front of FIFO queue and return

        Raise error if the queue is empty
        """

    @abstractmethod
    def first(self) -> Tuple[Queue, Any]:
        """Access (but do not remove) value at top of LIFO stack and return

        Raise Empty exception if queue is empty
        """


class ArrayQueue(Queue):
    """ArrayQueue implements a FIFO queue with a Python list as its underlying storage

    Using the Python list as underlying storage obscures the resizing operations and cost
    associated with dynamic arrays. However, as shown in chapter 1 the amortized cost of
    expanding and shrinking a dynamic array is O(n). Dynamic array constant factors may
    still be a bit fatter than a linked-list, depending on implementation details of each.
    Linked list implementations are likely to use more memory per element due to the need
    to hold references to the next (and previous in the case of a doubly-linked list) node
    """

    def __init__(self):
        self._arr = []

    def __len__(self) -> int:
        return len(self._arr)

    def is_empty(self) -> bool:
        return len(self._arr) == 0

    def enqueue(self, val: Any) -> ArrayQueue:
        self._arr.append(val)
        return self

    def dequeue(self) -> Tuple[ArrayQueue, Any]:
        if self.is_empty():
            raise Empty("queue is empty")
        val = self._arr[0]
        self._arr = self._arr[1:]
        return self, val

    def first(self) -> Tuple[ArrayQueue, Any]:
        if self.is_empty():
            raise Empty("queue is empty")
        val = self._arr[0]
        return self, val


class Empty(Exception):
    """Error attempting to access and element from an empty container type"""
