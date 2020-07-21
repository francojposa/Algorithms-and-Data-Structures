from abc import ABC, abstractmethod
from typing import Any, Iterable, Optional


class Tree(ABC):
    class Node(ABC):
        @abstractmethod
        def element(self) -> Any:
            """Return the element stored at this Node"""

        @abstractmethod
        def __eq__(self, other: "Node") -> bool:
            """Return True if other Node represents the same location in the Tree"""

        def __ne__(self, other: "Node") -> bool:
            """Return True if other Node does not represent the same location in the Tree"""
            return not (self == other)

    @abstractmethod
    def __len__(self) -> int:
        """Return the total number of Nodes in the Tree"""

    def is_empty(self) -> bool:
        """Return True if the tree is empty"""
        return len(self) == 0

    @abstractmethod
    def root(self) -> Optional[Node]:
        """Return Node representing the Tree's root or None if empty"""

    @abstractmethod
    def parent(self, node: Node) -> Optional[None]:
        """Return parent Node of node or None if node is root"""

    @abstractmethod
    def num_children(self, node: Node) -> int:
        """Return the count of nodes child Nodes"""

    @abstractmethod
    def children(self, node: Node) -> Iterable[Node]:
        """Return an Iterable of node's child Nodes"""

    def is_root(self, node: Node) -> bool:
        """Return True if node represents the root of the Tree"""
        return self.root() == node

    def is_leaf(self, node: Node) -> bool:
        """Return True if node does not have child Nodes"""
        return self.num_children(node) == 0

    def depth(self, node: Node) -> int:
        """Return the number of levels separating node from the root Node of the Tree"""
        if self.is_root(node):
            return 0
        return 1 + self.depth(self.parent(node))

    def height(self, node: Optional[Node] = None) -> int:
        """Return the height of the subtree rooted at node

        If node is None, return height of full Tree"
        """
        if node is None:
            node = self.root()
        return self._height(node)

    def _height(self, node: Node) -> int:
        """Return the height of the subtree rooted at node"""
        if self.is_leaf(node):
            return 0
        return 1 + max(self.height(c) for c in self.children(node))
