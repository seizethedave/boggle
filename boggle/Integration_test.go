package boggle

import (
   "bufio"
   "os"
   "testing"
)

const DoLargeFileTests = false

const sampleBoard = `apes
pies
tuba`

func TestLargeDictionaryFileWithBoard(t *testing.T) {
   if !DoLargeFileTests {
      t.Skip("skip large file test")
   }
   dictFile, err := os.Open("dict.txt")

   if err != nil {
      panic(err)
   }

   defer dictFile.Close()

   reader := bufio.NewReader(dictFile)
   dict := NewBoggleDictionaryFromReader(reader)
   board := NewBoardFromString(sampleBoard)
   words := board.ScanAll(dict)

   if !Contains(words, "apes") { t.Fail() }
   if !Contains(words, "pies") { t.Fail() }
   if !Contains(words, "tuba") { t.Fail() }
   if !Contains(words, "pit") { t.Fail() }
}

// Feed large file into dictionary, then re-scan the file and verify that the
// dictionary contains each word.
func TestLargeDictionaryExhaustive(t *testing.T) {
   if !DoLargeFileTests {
      t.Skip("skip large file tests")
   }

   dictFile, err := os.Open("dict.txt")

   if err != nil {
      panic(err)
   }

   defer dictFile.Close()

   reader := bufio.NewReader(dictFile)
   dict := NewBoggleDictionaryFromReader(reader)

   _, err = dictFile.Seek(0, 0)

   if err != nil {
      panic(err)
   }

   lineScanner := bufio.NewScanner(reader)
   for lineScanner.Scan() {
      word := lineScanner.Text()
      if !dict.Contains(word) {
         t.Errorf("Dict should contain %s but doesn't.", word)
      }
   }
}

