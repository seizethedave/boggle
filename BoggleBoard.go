package boggle

type BoggleNode struct {
   character rune
   connections []*BoggleNode
}

// Establishes a one-way connection.
func (node *BoggleNode) ConnectTo(otherNode *BoggleNode) {
   node.connections = append(node.connections, otherNode)
}

type BoggleBoard struct {
   nodes []*BoggleNode
}

func newBoardFromGrid(grid [][]rune) *BoggleBoard {

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
