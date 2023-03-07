package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursiveTreePreOrderMethod(t *testing.T) {

	tPart1 := Tree{
		Value: "Part 1",
		Children: []Tree{
			{
				Value: "Chapter 1",
				Children: []Tree{
					{Value: "1.1", Children: nil},
					{Value: "1.2", Children: nil},
				},
			},
			{
				Value: "Chapter 1",
				Children: []Tree{
					{Value: "1.1", Children: nil},
					{Value: "1.2", Children: nil},
				},
			},
		},
	}
	tableOfContents := []interface{}{}
	fn := func(t *Tree) interface{} {
		tableOfContents = append(tableOfContents, t.Value)
		return nil
	}
	tPart1.TraversePreOrder(fn)

	expectedTableOfContents := []interface{}{"Part 1", "Chapter 1", "1.1", "1.2", "Chapter 1", "1.1", "1.2"}
	assert.Equal(t, expectedTableOfContents, tableOfContents)
}
