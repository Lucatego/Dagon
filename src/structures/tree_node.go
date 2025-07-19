package structures

import (
	"errors"

	"dagon/src/syntax/valid_types"
)

type tree_node[T valid_types.Types] struct {
	data  T
	left  *tree_node[T]
	right *tree_node[T]
}

func CreateNode[T valid_types.Types](data T) *tree_node[T] {
	return &tree_node[T]{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (n *tree_node[T]) InsertData(data T) error {
	// If data is duplicated
	if data == n.data {
		return errors.New("the insertion of data failed, data already exists in the tree")
	} else if data < n.data { // If data is lower than the data in node
		if n.left != nil { // To prevent the access to null memory.
			return n.left.InsertData(data) // Recursion
		}
		n.left = CreateNode(data)
	} else if data > n.data {
		if n.right != nil {
			return n.right.InsertData(data)
		}
		n.right = CreateNode(data)
	}
	return nil
}
