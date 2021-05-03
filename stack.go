package queues

// Stack ... ist ein einfacher LIFO Stack
type Stack struct {
	nodes []*Node
	count int
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push ... fügt einen Node hinzu.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop ... entfernt den zuletzt hinzugefügten Node (LIFO) und liefert ihn zurück
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// Get ... liefert eine Liste von n Nodes, ohne sie zu entfernen
func (s *Stack) Get(n int, reverse bool) []*Node {
	if s.count == 0 {
		return nil
	}

	r := NewStack()
	count := 0
	// der neueste Node steht am Ende des Stacks
	if reverse == false {
		start := 0
		if n <= s.count-1 {
			start = s.count - n
		}
		for i := start; count <= n && i <= s.count-1; i++ {
			r.Push(NewNode(s.nodes[i].From, s.nodes[i].Content))
			count++
		}
		return r.nodes
	}
	// der neueste Node steht am Anfang des Stacks
	for i := s.count - 1; count < n && i >= 0; i-- {
		r.Push(NewNode(s.nodes[i].From, s.nodes[i].Content))
		count++
	}
	return r.nodes
}
