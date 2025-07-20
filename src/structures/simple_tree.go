package structures

import (
	"dagon/src/syntax/types"
	"errors"
	"fmt"
)

/*
	Simple_tree[T] struct:

It contains the head node from tree_node and the size of the tree. It also works
as a bridge for the functions in tree_node.
*/
type Simple_tree[T types.Types] struct {
	head *tree_node[T]
	size int64
}

func CreateTree[T types.Types]() *Simple_tree[T] {
	return &Simple_tree[T]{
		head: nil,
		size: 0,
	}
}

func (t *Simple_tree[T]) Add(data T) error {
	// First case
	if t.head == nil {
		t.head = createNode(data)
	} else {
		// Other cases
		err := t.head.insertData(data)
		if err != nil {
			return fmt.Errorf("error: Tree unable to insert new data [Unexpected Error: %s.]", err.Error())
		}
	}
	t.size++
	return nil
}

func (t *Simple_tree[T]) Find(data T) (bool, error) {
	if t.head == nil {
		return false, errors.New("error: The tree is empty [Usage error: Add data before using Find function]")
	}
	return t.head.find(data), nil
}
