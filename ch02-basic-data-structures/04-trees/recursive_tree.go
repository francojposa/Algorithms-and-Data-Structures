package trees

type Tree struct {
	Value    interface{}
	Children []Tree
}

type Visit func(t *Tree) interface{}

func (t *Tree) TraversePreOrder(visit Visit) {
	if visit != nil {
		visit(t)
	}
	for _, child := range t.Children {
		child.TraversePreOrder(visit)
	}
}

func TraversePreOrder(t *Tree, visit Visit) {
	if visit != nil {
		visit(t)
	}
	for _, child := range t.Children {
		child.TraversePreOrder(visit)
	}
}
