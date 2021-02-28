from typing import Any, Optional

from abc_tree import ITreeNode, IBinaryTree


class BinaryLinkedTree(IBinaryTree):
    class Node(ITreeNode):
        def __init__(
            self,
            container: "IBinaryTree",
            element: Any,
            parent: Optional["BinaryLinkedTree.Node"] = None,
            left_child: Optional["BinaryLinkedTree.Node"] = None,
            right_child: Optional["BinaryLinkedTree.Node"] = None,
        ):
            """Constructor should only be invoked by outer class, not directly by user"""
            self._container = container
            self._element = element
            self._parent = parent
            self._left_child = left_child
            self._right_child = right_child

        def __eq__(self, other: "BinaryLinkedTree.Node") -> bool:
            return all(
                [
                    type(self) == type(other),
                    self._element == other._element,
                    self._parent is other._parent,
                    self._left_child is other._left_child,
                    self._right_child is other._right_child,
                ]
            )

        def element(self) -> Any:
            return self._element

    def __init__(self):
        """Create an empty binary tree"""
        self._root: Optional["BinaryLinkedTree.Node"] = None
        self._size = 0

    def root(self) -> "BinaryLinkedTree.Node":
        return self._root

    def parent(
        self, node: "BinaryLinkedTree.Node"
    ) -> Optional["BinaryLinkedTree.Node"]:
        """Return parent Node of node or None if node is root"""
        node = self._validate(node)
        return node._parent

    def left(self, node: "BinaryLinkedTree.Node") -> Optional["BinaryLinkedTree.Node"]:
        """Return Node representing node's left child or None if no left child"""
        node = self._validate(node)
        return node._left

    def right(self, node: "BinaryLinkedTree.Node") -> Optional["BinaryLinkedTree.Node"]:
        """Return Node representing node's right child or None if no right child"""
        node = self._validate(node)
        return node._right

    def num_children(self, node: "BinaryLinkedTree.Node") -> int:
        # Could just count the length of children(), but am betting this implementation
        # is faster by not calling out to another method or creating a container type
        node = self._validate(node)
        num_children = 0
        if node._left_child is not None:
            num_children += 1
        if node._right_child is not None:
            num_children += 1

        return num_children

    def add_root(self, element: Any) -> "BinaryLinkedTree.Node":
        """Create root Node of an empty tree, containing element

        Return newly created root node
        """
        if self._root is not None:
            raise ValueError("non-empty tree; root node already exists")

        root_node = self.Node(self, element)
        self._root = root_node
        self._size = 1
        return root

    def add_left(
        self, node: "BinaryLinkedTree.Node", element: Any
    ) -> "BinaryLinkedTree.Node":
        """Create left child Node of node, containing element

        Return newly created left child node
        """
        parent_node = self._validate(node)
        if parent_node._left_child is not None:
            raise ValueError(f"left child already exists for node {parent_node}")

        left_child_node = self.Node(self, element, parent=parent_node)
        parent_node._left_child = left_child_node
        self._size += 1
        return left_child_node

    def add_right(
        self, node: "BinaryLinkedTree.Node", element: Any
    ) -> "BinaryLinkedTree.Node":
        """Create right child Node of node, containing element

        Return newly created right child node
        """
        parent_node = self._validate(node)
        if parent_node._right_child is not None:
            raise ValueError(f"right child already exists for node {parent_node}")

        right_child_node = self.Node(self, element, parent=parent_node)
        parent_node._right_child = right_child_node
        self._size += 1
        return right_child_node

    def _validate(self, node: "BinaryLinkedTree.Node") -> "BinaryLinkedTree.Node":
        if not isinstance(node, self.Node):
            raise TypeError(f"node {node} must be type BinaryLinkedTree.Node")
        if node._container is not self:
            raise ValueError(f"node {node} does not belong to this tree")
        if node._parent is node:
            # Convention for deleted nodes, aids garbage collection
            raise ValueError(
                f"node {node} has been removed from its tree and is no longer valid"
            )
