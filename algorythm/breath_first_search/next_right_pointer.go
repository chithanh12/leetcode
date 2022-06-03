package breath_first_search

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	nodesByLevel := make(map[int][]*Node, 0)
	populateMapNode(root, nodesByLevel, 0)
	for _, nodes := range nodesByLevel {
		for i := 0; i < len(nodes)-1; i++ {

			nodes[i].Next = nodes[i+1]
		}
	}

	return root
}

func populateMapNode(root *Node, nodesByLevel map[int][]*Node, level int) {
	if root == nil {
		return
	}

	if _, ok := nodesByLevel[level]; !ok {
		nodesByLevel[level] = make([]*Node, 0)
	}

	nodesByLevel[level] = append(nodesByLevel[level], root)

	populateMapNode(root.Left, nodesByLevel, level+1)
	populateMapNode(root.Right, nodesByLevel, level+1)
}
