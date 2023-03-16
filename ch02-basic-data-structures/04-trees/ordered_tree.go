package trees

type OrderedTree struct {
	value    any
	children []*OrderedTree
}
type OrderedTreeVisit func(tree *OrderedTree) error

func (ot *OrderedTree) Value() any {
	return ot.value
}

func (ot *OrderedTree) Children() []*OrderedTree {
	return ot.children
}

func (ot *OrderedTree) Height() int {
	if len(ot.children) == 0 {
		return 1
	}
	maxHeight := 0
	for _, child := range ot.children {
		childHeight := child.Height()
		if childHeight > maxHeight {
			maxHeight = childHeight
		}
	}
	return maxHeight + 1
}

func (ot *OrderedTree) TraversePreOrder(visit OrderedTreeVisit) error {
	if visit != nil {
		err := visit(ot)
		if err != nil {
			return err
		}
	}
	for _, child := range ot.Children() {
		err := child.TraversePreOrder(visit)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ot *OrderedTree) TraversePostOrder(visit OrderedTreeVisit) error {
	for _, child := range ot.Children() {
		err := child.TraversePostOrder(visit)
		if err != nil {
			return err
		}
	}
	if visit != nil {
		err := visit(ot)
		if err != nil {
			return err
		}
	}
	return nil
}
