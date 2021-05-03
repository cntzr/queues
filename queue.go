package queues

// Queue ... ist eine einfache FIFO Queue basierend auf einer zirkularen Liste
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// NewQueue ... liefert eine neue Queue mit einer initialen Länge zurück
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Push ... fügt einen Node hinzu
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop ... entfernt den zuerst hinzugefügten Node (FIFO) und liefert ihn zurück
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

// All ... liefert eine Liste aller Nodes, ohne sie zu entfernen
// der neueste Node steht am Ende der Queue!
func (q *Queue) All() []*Node {
	if q.count == 0 {
		return nil
	}
	return q.nodes
}
