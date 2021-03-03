from typing import List
from recursive_tree import Tree, traverse_preorder


def test_traverse_preorder_method():
    t1 = Tree(value="1", children=[])
    t1a = Tree(value="1a", children=[])
    t1b = Tree(value="1b", children=[])
    t1.children = [t1a, t1b]

    sections: List[str] = []

    t1.traverse_preorder(lambda t: sections.append(t.value))

    assert sections == ["1", "1a", "1b"]


def test_traverse_preorder_function():
    t1 = Tree(value="1", children=[])
    t1a = Tree(value="1a", children=[])
    t1b = Tree(value="1b", children=[])
    t1.children = [t1a, t1b]

    sections: List[str] = []

    traverse_preorder(t1, lambda t: sections.append(t.value))

    assert sections == ["1", "1a", "1b"]
