package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree_TraverseEuler(t *testing.T) {
	// expression tree for ((((3 + 1) * 3)/((9 − 5) + 2)) − ((3 * (7 − 4)) + 6))
	// the constructed expression string will be a valid python expression
	// the expression evaluates to -13
	expressionTree := &BinaryTree{
		value: "-",
		leftChild: &BinaryTree{
			value: "/",
			leftChild: &BinaryTree{
				value: "*",
				leftChild: &BinaryTree{
					value: "+",
					leftChild: &BinaryTree{
						value: "3",
					},
					rightChild: &BinaryTree{
						value: "1",
					},
				},
				rightChild: &BinaryTree{
					value:      "3",
					leftChild:  nil,
					rightChild: nil,
				},
			},
			rightChild: &BinaryTree{
				value: "+",
				leftChild: &BinaryTree{
					value: "-",
					leftChild: &BinaryTree{
						value: "9",
					},
					rightChild: &BinaryTree{
						value: "5",
					},
				},
				rightChild: &BinaryTree{
					value: "2",
				},
			},
		},
		rightChild: &BinaryTree{
			value: "+",
			leftChild: &BinaryTree{
				value: "*",
				leftChild: &BinaryTree{
					value: "3",
				},
				rightChild: &BinaryTree{
					value: "-",
					leftChild: &BinaryTree{
						value: "7",
					},
					rightChild: &BinaryTree{
						value: "4",
					},
				},
			},
			rightChild: &BinaryTree{
				value: "6",
			},
		},
	}
	expression := ""
	leftVisit := func(tree *BinaryTree) error {
		if tree.LeftChild() != nil && tree.RightChild() != nil {
			expression += "("
		}
		return nil
	}
	belowVisit := func(tree *BinaryTree) error {
		if tree.LeftChild() != nil || tree.RightChild() != nil {
			// not external; is an operand; pad with spaces
			expression += " " + tree.Value().(string) + " "
		} else {
			// external; is a value; do not pad
			expression += tree.Value().(string)
		}
		return nil
	}
	rightVisit := func(tree *BinaryTree) error {
		if tree.LeftChild() != nil && tree.RightChild() != nil {
			expression += ")"
		}
		return nil
	}
	err := expressionTree.TraverseEuler(leftVisit, belowVisit, rightVisit)
	assert.Nil(t, err)

	assert.Equal(t, "((((3 + 1) * 3) / ((9 - 5) + 2)) - ((3 * (7 - 4)) + 6))", expression)
}
