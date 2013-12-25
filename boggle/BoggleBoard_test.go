package boggle

import (
   "testing"
   "fmt"
)

var _ = fmt.Printf

func Contains(words []string, word string) bool {
   for _, w := range words {
      if w == word {
         return true
      }
   }
   return false
}


const board1 = `abc
def
ghi`

func TestCreateFromString(t *testing.T) {
   board := NewBoardFromString(board1)

   if len(board.nodes) != 9 {
      t.Errorf("Should be 9 nodes. got %d instead. (%+v)", len(board.nodes),
       board.nodes)
   }

   words := board.ScanAll(nil)

   if !Contains(words, "abcfedghi") {
      t.Errorf("expected word not found among %d words", len(words))
      t.Logf("%v", words)
   }
   t.Logf("%+v", board.nodes[0])
}

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

   words := make([]string, 0)

   board.Scan(func(word string) {
      words = append(words, word)
   }, nil)

   if !Contains(words, "a") { t.Fail() }
   if !Contains(words, "ab") { t.Fail() }
   if !Contains(words, "abc") { t.Fail() }
   if !Contains(words, "abcd") { t.Fail() }
   if !Contains(words, "abdc") { t.Fail() }
   if !Contains(words, "b") { t.Fail() }
   if !Contains(words, "bcad") { t.Fail() }
   if !Contains(words, "bcda") { t.Fail() }
}

func TestScan2x2WithDict(t *testing.T) {
   board := NewBoardFromGrid([][]rune {
     { 'a', 'b' },
     { 'c', 'd' },
   })

   dict := NewBoggleDictionaryWithWords(
    "ab", "ad", "bad", "a", "dab",
    "drab", "brad", "crab")

   words := board.ScanAll(dict)

   // Words that should be in there:
   if !Contains(words, "a") { t.Errorf("a") }
   if !Contains(words, "ab") { t.Errorf("ab") }
   if !Contains(words, "ad") { t.Errorf("ad") }
   if !Contains(words, "bad") { t.Errorf("bad") }
   if !Contains(words, "dab") { t.Errorf("dab") }

   // Words that should not:
   if Contains(words, "drab") { t.Errorf("drab") }
   if Contains(words, "brad") { t.Errorf("brad") }
   if Contains(words, "crab") { t.Errorf("crab") }
}
