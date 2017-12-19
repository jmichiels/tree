package tree

import (
	"fmt"
	"testing"
)

type TestNode struct {
	ID     int
	Name   string
	Parent int
}

func (node TestNode) String() string {
	return fmt.Sprintf("%s (%d)", node.Name, node.ID)
}

type TestTree []TestNode

func (tree TestTree) RootNodes() (nodes []Node) {
	return tree.ChildrenNodes(TestNode{})
}

func (tree TestTree) ChildrenNodes(parent Node) (nodes []Node) {
	for _, node := range tree {
		if node.Parent == parent.(TestNode).ID {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func Test(t *testing.T) {

	tree := TestTree{
		TestNode{
			ID:   1,
			Name: "First Root Node",
		},
		TestNode{
			ID:     11,
			Name:   "A child node",
			Parent: 1,
		},
		TestNode{
			ID:     12,
			Name:   "Another child node",
			Parent: 1,
		},
		TestNode{
			ID:     121,
			Name:   "A child child node",
			Parent: 12,
		},
		TestNode{
			ID:     1211,
			Name:   "A child child child node",
			Parent: 121,
		},
		TestNode{
			ID:     13,
			Name:   "Yet another one",
			Parent: 1,
		},
		TestNode{
			ID:   2,
			Name: "Second Root Node",
		},
		TestNode{
			ID:     21,
			Name:   "Yet another child node",
			Parent: 2,
		},
	}

	ascii := `First Root Node (1)
├── A child node (11)
├── Another child node (12)
│   └── A child child node (121)
│       └── A child child child node (1211)
└── Yet another one (13)
Second Root Node (2)
└── Yet another child node (21)
`

	if String(tree) != ascii {
		t.Error("the generated tree does not match the expected one")
	}
}
