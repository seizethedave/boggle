package boggle

type BoggleNode struct {
   character rune
   connections []*BoggleNode
}

// Establishes a one-way connection.
func (node *BoggleNode) ConnectTo(otherNode *BoggleNode) {
   node.connections = append(node.connections, otherNode)
}
