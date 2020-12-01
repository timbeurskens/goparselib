package goparselib

import (
	"errors"
	"fmt"
	"io"
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
func (n *Node) Populate(str string) {
	if n.Children == nil || len(n.Children) == 0 {
		n.Contents = str[n.Start : n.Size+n.Start]
		return
	}

	for i := range n.Children {
		n.Children[i].Populate(str)
	}
}

// Reduce creates a reduced copy of the original tree by removing all nodes with type in symbols
func (n Node) Reduce(symbols ...Symbol) (Node, error) {
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
