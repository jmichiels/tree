# Tree
This golang package creates a nice tree representation of your data:
```text
First Root Node (1)
├── A child node (11)
├── Another child node (12)
│   └── A child child node (121)
│       └── A child child child node (1211)
└── Yet another one (13)
Second Root Node (2)
└── Yet another child node (21)
```
## Usage
Just create a type for your tree which implements the [`tree.Tree`](tree.go#L9) interface (see for example the [`TestTree`](tree_test.go#L18) type used for testing):
```go
type Tree interface {

	// Returns the top level nodes.
	RootNodes() []Node

	// Returns the children of the provided node.
	ChildrenNodes(Node) []Node
}
```
Then, provide an instance of this type to `tree.String` or `tree.Write`, depending if you prefer the whole tree as a string or progressively written to the output via an `io.Writer`.