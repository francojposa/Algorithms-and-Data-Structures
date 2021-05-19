from typing import List
from recursive_tree import Tree, traverse_preorder


def test_traverse_preorder_method():
    t_section = Tree(value="section", children=[])
    t1 = Tree(value="1", children=[])
    t1a = Tree(value="1a", children=[])
    t1b = Tree(value="1b", children=[])
    t2 = Tree(value="2", children=[])
    t2a = Tree(value="2a", children=[])
    t1.children = [t1a, t1b]
    t2.children = [t2a]
    t_section.children = [t1, t2]

    table_of_contents: List[str] = []

    t_section.traverse_preorder(lambda t: table_of_contents.append(t.value))

    assert table_of_contents == ["section", "1", "1a", "1b", "2", "2a"]


def test_traverse_preorder_function():
    t_section = Tree(
        value="section",
        children=[
            Tree(
                value="1",
                children=[Tree(value="1a", children=[]), Tree(value="1b", children=[])],
            ),
            Tree(value="2", children=[Tree(value="2a", children=[])]),
        ],
    )
    table_of_contents: List[str] = []

    traverse_preorder(t_section, lambda t: table_of_contents.append(t.value))

    assert table_of_contents == ["section", "1", "1a", "1b", "2", "2a"]

