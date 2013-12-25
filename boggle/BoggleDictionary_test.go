package boggle

import (
   "bufio"
   "os"
   "testing"
   "strings"
)

func TestCreateFromWords(t *testing.T) {
   dict := NewBoggleDictionaryWithWords("ape", "aardvark", "apple")
   if !dict.Contains("ape") || !dict.Contains("aardvark") ||
    !dict.Contains("apple") {
      t.Fail()
   }

   if dict.Contains("a") {
      t.Errorf("Should not contain a")
   }
}

func TestCreateAbAdIssue(t *testing.T) {
   // There's an issue in the dictionary with a/ad/ab.
   dict := NewBoggleDictionaryWithWords("a", "ad", "ab", "ape",
    "pizza", "apes", "bugle", "bug")
   if !dict.Contains("a") || !dict.Contains("ad") || !dict.Contains("ab") {
      t.Errorf("doesn't contain a/ad/ab")
   }
}

const DummyFile = `ape
chicken
whale
donkey
tuna
blubber
quail
jinkle
`

func TestCreateFromReader(t *testing.T) {
   reader := strings.NewReader(DummyFile)
   dict := NewBoggleDictionaryFromReader(reader)

   if !dict.Contains("tuna") {
      t.Fail()
   }

   if dict.Contains("lemur") {
      t.Fail()
   }
}

func TestCreateFromActualFile(t *testing.T) {
   dictFile, err := os.Open("sample.txt")

   if err != nil {
      panic(err)
   }

   defer dictFile.Close()

   reader := bufio.NewReader(dictFile)
   dict := NewBoggleDictionaryFromReader(reader)
   if !dict.Contains("aardvark") { t.Fail() }
}

