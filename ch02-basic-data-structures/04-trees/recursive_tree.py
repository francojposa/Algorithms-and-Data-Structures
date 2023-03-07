from __future__ import annotations
from typing import Any, Callable, Iterable, Optional

import pydantic


class Tree(pydantic.BaseModel):
    value: Any
    children: Iterable[Tree]

    def traverse_preorder(self, visit: Optional[Callable[[Tree], Any]] = None):
        if visit:
            visit(self)
        for child in self.children:
            child.traverse_preorder(visit)

    def traverse_postorder(self, visit: Optional[Callable[[Tree], Any]] = None):
        for child in self.children:
            child.traverse_postorder(visit)
        if visit:
            visit(self)


Tree.update_forward_refs()


def traverse_preorder(tree: Tree, visit: Optional[Callable[[Tree], Any]] = None):
    if visit:
        visit(tree)
    for child in tree.children:
        traverse_preorder(child, visit)
