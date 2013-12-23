package boggle

import (
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
