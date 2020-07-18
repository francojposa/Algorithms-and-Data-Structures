import pytest

from array_stack import ArrayStack, Empty


def test_array_stack():
    array_stack = ArrayStack()

    # Assert that we start with an empty stack
    assert array_stack.is_empty()
    with pytest.raises(Empty):
        array_stack.top()

    # Push a bunch of elements onto the stack
    for i in range(10):
        array_stack.push(i)
        array_stack, top = array_stack.top()
        assert top == i

    # Pop all the elements off
    for i in range(len(array_stack) - 1, -1, -1):
        array_stack, val = array_stack.pop()
        assert val == i

    # Assert we are back to an empty stack
    assert array_stack.is_empty()
    with pytest.raises(Empty):
        array_stack.pop()
