package boggle

type BoggleNode struct {
   character rune
   connections []*BoggleNode

   // These are used during traversal:
   discovered bool
   visitParent *BoggleNode
   visitStack []*BoggleNode
}

// Establishes a one-way connection.
func (node *BoggleNode) ConnectTo(otherNode *BoggleNode) {
   node.connections = append(node.connections, otherNode)
}

func (node *BoggleNode) Visited() bool {
   return node.visitStack != nil
}

func (node *BoggleNode) Explore() {
   node.visitStack = make([]*BoggleNode, 0)

   for _, connection := range node.connections {
      if (!connection.discovered) {
         node.visitStack = append(node.visitStack, connection)
      }
   }
}

func (node *BoggleNode) Abandon() {
   node.visitStack = nil
}

