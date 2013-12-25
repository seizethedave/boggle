package boggle

import (
   "strings"
)

type BoggleBoard struct {
   nodes []*BoggleNode
}

type FoundWordFunc func(string)

func NewBoardFromString(grid string) *BoggleBoard {
   // Turn encoded grid string into NxM nested array and pass it to
   // NewBoardFromGrid.

   runeGrid := make([][]rune, 0)

   for _, row := range strings.Split(grid, "\n") {
      runeGrid = append(runeGrid, []rune(row))
   }

   return NewBoardFromGrid(runeGrid)
}

func NewBoardFromGrid(grid [][]rune) *BoggleBoard {
   height := len(grid)
   width := len(grid[0])

   var nodeSlice []*BoggleNode;
   nodeGrid := make([][]*BoggleNode, height)

   for y, row := range grid {
      nodeGrid[y] = make([]*BoggleNode, len(row))

      for x, char := range row {
         newNode := BoggleNode { character: char }
         nodeSlice = append(nodeSlice, &newNode)
         nodeGrid[y][x] = &newNode
      }
   }

   var thisNode *BoggleNode
   var notRightCol, notLeftCol, notTopRow, notBottomRow bool

   for y := 0; y < height; y++ {
      notTopRow = y > 0
      notBottomRow = y < height - 1

      for x := 0; x < width; x++ {
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

// Implements the common case where you just want to get an array of all words.
func (board *BoggleBoard) ScanAll(dict *BoggleDictionary) (words []string) {
   words = make([]string, 0)

   board.Scan(func(word string) {
      words = append(words, word)
   }, dict)

   return
}

// Calls foundFunc with each word discovered in the board.
func (board *BoggleBoard) Scan(foundFunc FoundWordFunc, dict *BoggleDictionary) {

   var navigator DictionaryNavigator
   if dict != nil {
      navigator = NewBoggleDictionaryNavigator(dict.root)
   } else {
      navigator = &DumbDictionaryNavigator { }
   }

   var ledger []rune

   for _, root := range board.nodes {
      node := root

      for node != nil {

         // Visited() will return true during a backtrack. We know this path has
         // been OK'd by the dictionary sometime in the past so we don't need to
         // ask again.

         if !node.Visited() {

            if navigator.TryPush(node.character) {
               node.discovered = true
               node.Visit()

               ledger = append(ledger, node.character)

               if navigator.EndOfWord() {
                  foundFunc(string(ledger))
               }
            }
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
            node.discovered = false

            if orphan.Visited() {
               navigator.Pop()
               ledger = ledger[:len(ledger) - 1]
               orphan.Unvisit()
            }

            node = orphan.visitParent
            orphan.visitParent = nil
         }
      }
   }
}
