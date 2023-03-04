import pytest

from array_queue import ArrayQueue, Empty


def test_array_queue():
    array_queue = ArrayQueue()

    # Assert that we start with an empty queue
    assert array_queue.is_empty()
    with pytest.raises(Empty):
        array_queue.first()

    # Enqueue a bunch of elements into the back of the queue
    for i in range(10):
        array_queue.enqueue(i)
        array_queue, first = array_queue.first()
        assert first == 0
        assert len(array_queue) == i + 1

    # Dequeue a bunch of elements from the front of the queue
    for i in range(10):
        array_queue, val = array_queue.dequeue()
        assert val == i
        assert len(array_queue) == 10 - 1 - i

    # Assert we are back to an empty queue
    assert array_queue.is_empty()
    with pytest.raises(Empty):
        array_queue.dequeue()
