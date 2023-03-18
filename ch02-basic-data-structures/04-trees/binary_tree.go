package trees

type BinaryTree struct {
	value      any
	leftChild  *BinaryTree
	rightChild *BinaryTree
}

type BinaryTreeVisit func(tree *BinaryTree) error

func (bt *BinaryTree) Value() any {
	return bt.value
}

func (bt *BinaryTree) LeftChild() *BinaryTree {
	return bt.leftChild
}

func (bt *BinaryTree) RightChild() *BinaryTree {
	return bt.rightChild
}

func (bt *BinaryTree) TraverseEuler(leftVisit, belowVisit, rightVisit BinaryTreeVisit) error {
	if leftVisit != nil {
		err := leftVisit(bt)
		if err != nil {
			return err
		}
	}

	if leftChild := bt.LeftChild(); leftChild != nil {
		err := leftChild.TraverseEuler(leftVisit, belowVisit, rightVisit)
		if err != nil {
			return err
		}
	}
	if belowVisit != nil {
		err := belowVisit(bt)
		if err != nil {
			return err
		}
	}

	if rightChild := bt.RightChild(); rightChild != nil {
		err := rightChild.TraverseEuler(leftVisit, belowVisit, rightVisit)
		if err != nil {
			return err
		}
	}
	if rightVisit != nil {
		err := rightVisit(bt)
		if err != nil {
			return err
		}
	}

	return nil
}
