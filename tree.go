package tree

import (
	"bytes"
	"fmt"
	"io"
)

type Tree interface {

	// Returns the top level nodes.
	RootNodes() []Node

	// Returns the children of the provided node.
	ChildrenNodes(Node) []Node
}

type Node fmt.Stringer

// String returns the whole tree as a string.
func String(tree Tree) string {
	var buffer bytes.Buffer
	Write(tree, &buffer)
	return buffer.String()
}

// Write writes the tree line by line to the specified writer.
func Write(tree Tree, writer io.Writer) error {
	return format(tree, tree.RootNodes(), true, "", writer)
}

const (
	dash       string = `├── `
	spacer     string = `│   `
	dashLast   string = `└── `
	spacerLast string = `    `
)

func format(tree Tree, nodes []Node, root bool, prefix string, output io.Writer) error {
	for idx, node := range nodes {
		lineBuffer := prefix
		childPrefix := prefix
		if !root {
			switch idx {
			case (len(nodes) - 1):
				// Last of the subnodes.
				lineBuffer += dashLast
				childPrefix += spacerLast

			default:
				lineBuffer += dash
				childPrefix += spacer
			}
		}
		lineBuffer += node.String() + "\n"
		// Write node string representation to output.
		if _, err := output.Write([]byte(lineBuffer)); err != nil {
			return err
		}
		// Write childre lines to output.
		if err := format(tree, tree.ChildrenNodes(node), false, childPrefix, output); err != nil {
			return err
		}
	}
	return nil
}
