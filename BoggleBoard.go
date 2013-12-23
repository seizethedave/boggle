package boggle

type BoggleBoard struct {
   nodes []*BoggleNode
}

type FoundWordFunc func(string)

func NewBoardFromGrid(grid [][]rune) *BoggleBoard {
   height := len(grid)
   width := len(grid[0])

   var nodeSlice []*BoggleNode;
   nodeGrid := make([][]*BoggleNode, height)

   var x, y int

   for _, row := range grid {
      nodeGrid[y] = make([]*BoggleNode, len(row))

      x = 0

      for _, char := range row {
         newNode := BoggleNode { character: char }
         nodeSlice = append(nodeSlice, &newNode)
         nodeGrid[y][x] = &newNode
         x++
      }

      y++
   }

   var thisNode *BoggleNode
   var notRightCol, notLeftCol, notTopRow, notBottomRow bool

   for y = 0; y < height; y++ {
      notTopRow = y > 0
      notBottomRow = y < height - 1

      for x = 0; x < width; x++ {
         thisNode = nodeGrid[y][x]
         notLeftCol = x > 0
         notRightCol = x < width - 1

         if notTopRow {
            // Up.
            thisNode.ConnectTo(nodeGrid[y - 1][x])
         }

         if notRightCol {
            if notTopRow {
               // Up/right.
               thisNode.ConnectTo(nodeGrid[y - 1][x + 1])
            }

            // Right.
            thisNode.ConnectTo(nodeGrid[y][x + 1])

            if notBottomRow {
               // Down/right.
               thisNode.ConnectTo(nodeGrid[y + 1][x + 1])
            }
         }

         if notLeftCol {
            if notTopRow {
               // Up/left.
               thisNode.ConnectTo(nodeGrid[y - 1][x - 1])
            }

            // Left.
            thisNode.ConnectTo(nodeGrid[y][x - 1])

            if notBottomRow {
               // Down/left.
               thisNode.ConnectTo(nodeGrid[y + 1][x - 1])
            }
         }

         if notBottomRow {
            // Down.
            thisNode.ConnectTo(nodeGrid[y + 1][x])
         }
      }
   }

   return &BoggleBoard { nodes: nodeSlice };
}

func (board *BoggleBoard) Scan(foundFunc FoundWordFunc) {
   var ledger []rune

   for _, root := range board.nodes {
      node := root

      for node != nil {
         pursueNode := (node.Visited() || true /* todo: navigator */)

         if pursueNode && !node.Visited() {
            node.discovered = true
            node.Visit()

            ledger = append(ledger, node.character)

            // todo: do if end of word.
            foundFunc(string(ledger))
         }

         if node.visitStack != nil && len(node.visitStack) > 0 {
            // There are unvisited adjacent nodes. Continue.
            var unseenNeighbor *BoggleNode
            unseenNeighbor, node.visitStack =
             node.visitStack[len(node.visitStack) - 1],
             node.visitStack[:len(node.visitStack) - 1]

            parent := node
            node = unseenNeighbor
            node.visitParent = parent
         } else {
            // Already visited all worthy edges from this node. Backtrack.
            orphan := node
            orphan.discovered = false

            if orphan.Visited() {
               // todo: pop navigator
               ledger = ledger[:len(ledger) - 1]
               orphan.Unvisit()
            }

            node = orphan.visitParent
            orphan.visitParent = nil
         }
      }
   }
}
