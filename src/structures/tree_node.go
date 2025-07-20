package structures

import (
	"dagon/src/syntax/types"
	"errors"
)

type tree_node[T types.Types] struct {
	data  T
	left  *tree_node[T]
	right *tree_node[T]
}

func createNode[T types.Types](data T) *tree_node[T] {
	return &tree_node[T]{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (n *tree_node[T]) insertData(data T) error {
	// If data is duplicated
	if data == n.data {
		return errors.New("data already exists in the tree")
	} else if data < n.data { // If data is lower than the data in node
		if n.left != nil { // To prevent the access to null memory.
			return n.left.insertData(data) // Recursion
		}
		n.left = createNode(data)
	} else if data > n.data {
		if n.right != nil {
			return n.right.insertData(data)
		}
		n.right = createNode(data)
	}
	return nil
}

func (n *tree_node[T]) find(data T) bool {
	// It's found
	if n.data == data {
		return true
	}
	// Look in the tree
	if n.right != nil && data > n.data { // Only if the node exists
		return n.right.find(data)
	} else if n.left != nil && data < n.data { // Only if the node exists
		return n.left.find(data)
	}
	// Base case
	return false
}
