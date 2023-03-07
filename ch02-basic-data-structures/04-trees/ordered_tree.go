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
