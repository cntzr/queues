package queues

type (
	// Node ... beschreibt einen einzelnen Eintrag in einer Queue oder auf einem Stack
	Node struct {
		From    string
		Content string
	}
)

// NewNode ... generiert einen neuen Node mit dem übergebenen Inhalt
func NewNode(from, content string) *Node {
	return &Node{
		From:    from,
		Content: content,
	}
}

func (n *Node) getFrom() string {
	return n.From
}

func (n *Node) getContent() string {
	return n.Content
}
