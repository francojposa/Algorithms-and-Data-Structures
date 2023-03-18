package trees

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedTree_TraversePreOrder(t *testing.T) {
	var tableOfContentsTree = &OrderedTree{
		value: nil,
		children: []*OrderedTree{
			{
				value: "Chapter 1",
				children: []*OrderedTree{
					{
						value:    "1.1",
						children: nil,
					},
					{
						value:    "1.2",
						children: nil,
					},
				},
			},
			{
				value: "Chapter 2",
				children: []*OrderedTree{
					{
						value:    "2.1",
						children: nil,
					},
				},
			},
		},
	}

	assert.Equal(t, tableOfContentsTree.Height(), 3)

	var tableOfContents []string
	visit := func(tree *OrderedTree) error {
		val := tree.Value()
		if val == nil {
			return nil
		}
		str, ok := val.(string)
		if !ok {
			return errors.New("tree value is not a string")
		}
		tableOfContents = append(tableOfContents, str)
		return nil
	}

	err := tableOfContentsTree.TraversePreOrder(visit)
	assert.Nil(t, err)

	expectedTableOfContents := []string{"Chapter 1", "1.1", "1.2", "Chapter 2", "2.1"}
	assert.Equal(t, expectedTableOfContents, tableOfContents)
}

func TestOrderedTree_TraversePostOrder(t *testing.T) {
	var reversePolishNotationTree = &OrderedTree{
		value: "-",
		children: []*OrderedTree{
			{
				value: "*",
				children: []*OrderedTree{
					{
						value: "+",
						children: []*OrderedTree{
							{
								value:    "2",
								children: nil,
							},
							{
								value:    "3",
								children: nil,
							},
						},
					},
					{
						value:    "y",
						children: nil,
					},
				},
			},
			{
				value:    "2",
				children: nil,
			},
		},
	}

	assert.Equal(t, reversePolishNotationTree.Height(), 4)

	var reversePolishNotation []string
	visit := func(tree *OrderedTree) error {
		val := tree.Value()
		if val == nil {
			return nil
		}
		str, ok := val.(string)
		if !ok {
			return errors.New("tree value is not a string")
		}
		reversePolishNotation = append(reversePolishNotation, str)
		return nil
	}

	err := reversePolishNotationTree.TraversePostOrder(visit)
	assert.Nil(t, err)

	expectedReversePolishNotation := []string{"2", "3", "+", "y", "*", "2", "-"}
	assert.Equal(t, expectedReversePolishNotation, reversePolishNotation)
}
