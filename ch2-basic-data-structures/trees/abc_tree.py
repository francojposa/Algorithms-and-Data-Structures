from abc import ABC, abstractmethod
from typing import Any, Iterable, List, Optional


class ITreeNode(ABC):
    @abstractmethod
    def element(self) -> Any:
        """Return the element stored at this Node"""

    @abstractmethod
    def __eq__(self, other: "ITreeNode") -> bool:
        """Return True if other Node represents the same location in the Tree"""

    def __ne__(self, other: "ITreeNode") -> bool:
        """Return True if other Node does not represent the same location in the Tree"""
        return not (self == other)


class ITree(ABC):
    @abstractmethod
    def __len__(self) -> int:
        """Return the total number of Nodes in the Tree"""

    def is_empty(self) -> bool:
        """Return True if the tree is empty"""
        return len(self) == 0

    @abstractmethod
    def nodes(self) -> Iterable[ITreeNode]:
        """Return an Iterable of all Nodes in the Tree"""

    def elements(self) -> Iterable[Any]:
        """Return an Iterable of all elements contained in Nodes in the Tree"""
        nodes = self.nodes()
        return [node.element() for node in nodes]

    @abstractmethod
    def root(self) -> Optional[ITreeNode]:
        """Return Node representing the Tree's root or None if empty"""

    @abstractmethod
    def parent(self, node: ITreeNode) -> Optional[None]:
        """Return parent Node of node or None if node is root"""

    @abstractmethod
    def num_children(self, node: ITreeNode) -> int:
        """Return the count of nodes child Nodes"""

    @abstractmethod
    def children(self, node: ITreeNode) -> Iterable[Optional[ITreeNode]]:
        """Return an Iterable of node's child Nodes

        Returning None to represent a missing child Node may be appropriate
        for tree structures defined by a fixed number of children.
        Ex in the case of a binary tree node with no left child, returning
        [None, some_node] is more informative than skipping the missing child
        and just returning [some_node] - is some_node the left or right child?
        """

    def is_root(self, node: ITreeNode) -> bool:
        """Return True if node represents the root of the Tree"""
        return self.root() == node

    def is_leaf(self, node: ITreeNode) -> bool:
        """Return True if node does not have child Nodes"""
        return self.num_children(node) == 0

    def depth(self, node: ITreeNode) -> int:
        """Return the number of levels separating node from the root Node of the Tree"""
        if self.is_root(node):
            return 0
        return 1 + self.depth(self.parent(node))

    def height(self, node: Optional[ITreeNode] = None) -> int:
        """Return the height of the subtree rooted at node

        If node is None, return height of full Tree"
        """
        if node is None:
            node = self.root()
        return self._height(node)

    def _height(self, node: ITreeNode) -> int:
        """Return the height of the subtree rooted at node"""
        if self.is_leaf(node):
            return 0
        return 1 + max(self.height(c) for c in self.children(node))


class IBinaryTree(ITree):
    @abstractmethod
    def left(self, node: ITreeNode) -> Optional[ITreeNode]:
        """Return Node representing node's left child or None if no left child"""

    @abstractmethod
    def right(self, node: ITreeNode) -> Optional[ITreeNode]:
        """Return Node representing node's right child or None if no right child"""

    def sibling(self, node: ITreeNode) -> Optional[ITreeNode]:
        """Return Node representing node's sibling or None if no sibling"""
        parent = self.parent(node)
        if parent is None:
            return None
        parent_left_child = self.left(parent)
        if node == parent_left_child:
            return self.right(parent)
        return parent_left_child

    def children(self, node: ITreeNode) -> List[Optional[ITreeNode]]:
        return [self.left(node), self.right(node)]
