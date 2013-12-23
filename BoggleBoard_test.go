package boggle

import (
   "testing"
   "fmt"
)

var _ = fmt.Printf

func TestCreate1x1(t *testing.T) {
   board := NewBoardFromGrid([][]rune { { 'c' } })

   if len(board.nodes) != 1 {
      t.Fail()
   }
}

func TestLength2x2(t *testing.T) {
   board := NewBoardFromGrid([][]rune {
     { 'a', 'b' },
     { 'c', 'd' },
   })

   if len(board.nodes) != 4 {
     t.Fail()
   }
}

func TestConnections3x3(t *testing.T) {
   board := NewBoardFromGrid([][]rune {
     { 'a', 'b', 'x' },
     { 'c', 'd', 'y' },
     { 'e', 'r', 't' },
   })

   if len(board.nodes[0].connections) != 3 {
      t.Fail()
   }

   if len(board.nodes[1].connections) != 5 {
      t.Fail()
   }

   if len(board.nodes[4].connections) != 8 {
      t.Fail()
   }
}

func TestScan2x2(t *testing.T) {
   board := NewBoardFromGrid([][]rune {
     { 'a', 'b' },
     { 'c', 'd' },
   })

   board.Scan(func(word string) {
      // fmt.Println("Yar", word)
   })
}
