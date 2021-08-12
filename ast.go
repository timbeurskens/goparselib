package goparselib

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

func inSymbols(s Symbol, symbols []Symbol) bool {
	for i := range symbols {
		if symbols[i] == s {
			return true
		}
	}
	return false
}

// Populate maps a given string to the node structure by placing substrings in the corresponding leaf nodes
// Deprecated: Newer versions of the parser automatically populate the parse tree
func (n *Node) Populate(str string) {
	if n.Children == nil || len(n.Children) == 0 {
		n.Contents = str[n.Start : n.Size+n.Start]
		return
	}

	for i := range n.Children {
		n.Children[i].Populate(str)
	}
}

// Collapse maps a nil-reduced, recurrent tree to a list under the given condition
func (n Node) Collapse(selectFn func(Node) (Node, bool), walkFn func(Node) bool) []Node {
	result := make([]Node, 0)

	n.Walk(func(node Node) bool {
		if n, ok := selectFn(node); ok {
			result = append(result, n)
		}
		return walkFn(node)
	})

	return result
}

func (n Node) CollapseType(symbol Symbol) ([]Node, error) {
	return n.Collapse(func(node Node) (Node, bool) {
		if reflect.DeepEqual(node.Type, symbol) {
			return node, true
		}
		return Node{}, false
	}, func(node Node) bool {
		return true
	}), nil
}

// NilReduce ensures no nil symbols exist in the tree
func (n Node) NilReduce() (Node, error) {
	return n.Reduce(nil)
}

// Reduce creates a reduced copy of the original tree by removing all nodes with type in symbols
func (n Node) Reduce(symbols ...Symbol) (Node, error) {
	if symbols == nil || len(symbols) == 0 {
		return Node{}, errors.New("set of layout symbols is empty or nonexistent")
	}

	if inSymbols(n.Type, symbols) {
		return Node{}, errors.New("root node is not in the reduced set")
	}

	if n.Children == nil {
		return n, nil
	}

	result := make([]Node, 0, len(n.Children))
	for i := range n.Children {
		if inSymbols(n.Children[i].Type, symbols) {
			continue
		}
		if r, err := n.Children[i].Reduce(symbols...); err != nil {
			continue
		} else {
			result = append(result, r)
		}
	}
	return Node{
		Start:    n.Start,
		Size:     n.Size,
		Contents: n.Contents,
		Type:     n.Type,
		Children: result,
	}, nil
}

func (n Node) String() string {
	return fmt.Sprintf("(\n%s[%d,%d]:%s:%s\n)\n", n.Type, n.Start, n.Size, n.Contents, n.Children)
}

func (n Node) Output(writer io.Writer) error {
	if _, err := writer.Write([]byte(n.Contents)); err != nil {
		return err
	}

	if n.Children == nil {
		return nil
	}

	for i := range n.Children {
		if err := n.Children[i].Output(writer); err != nil {
			return err
		}
	}

	return nil
}

func (n Node) FindType(symbol Symbol) (Node, error) {
	return n.Find(func(node Node) bool {
		return reflect.DeepEqual(node.Type, symbol)
	})
}

// Find returns the first node in which the provided condition holds
func (n Node) Find(condition func(Node) bool) (Node, error) {
	if condition(n) {
		return n, nil
	} else {

		if n.Children == nil {
			return Node{}, errors.New("not found")
		}

		for i := range n.Children {
			r2 := n.Children[i]
			if r3, err := r2.Find(condition); err == nil {
				return r3, nil
			}
		}

		return Node{}, errors.New("not found")
	}
}

func (n Node) Walk(continueCondition func(Node) bool) {
	if !continueCondition(n) {
		return
	} else if n.Children != nil {
		for c := range n.Children {
			n.Children[c].Walk(continueCondition)
		}
	}
}
