package boggle

import (
   "testing"
)

func TestInterfaces (t *testing.T) {
   // Assign implementers to var of interface type to make sure they're legit.

   var nav DictionaryNavigator
   nav = &BoggleDictionaryNavigator { }
   nav = &DumbDictionaryNavigator { }
   var _ = nav
}

func TestPushPop(t *testing.T) {
   dict := NewBoggleDictionaryWithWords(
    "a", "ad", "ab", "ape", "pizza", "apes", "bugle", "bug")
   navigator := NewBoggleDictionaryNavigator(dict.root)

   if !navigator.TryPush('a') { t.Errorf("push should work") }
   if !navigator.TryPush('p') { t.Errorf("push should work") }
   if !navigator.TryPush('e') { t.Errorf("push should work") }
   if !navigator.TryPush('s') { t.Errorf("push should work") }

   if navigator.TryPush('x') { t.Errorf("push should fail") }

   if !navigator.EndOfWord() { t.Errorf("apes is end of word.") }
   navigator.Pop()
   if !navigator.EndOfWord() { t.Errorf("ape is end of word.") }

   navigator.Pop()
   t.Logf("%c %+v", navigator.node.character, navigator.node)
   navigator.Pop()
   t.Logf("%c %+v", navigator.node.character, navigator.node)
   if !navigator.EndOfWord() { t.Errorf("a is end of word.") }

   if !navigator.TryPush('b') { t.Errorf("push should work") }
   t.Logf("%c %+v", navigator.node.character, navigator.node)
   if !navigator.EndOfWord() { t.Errorf("ab is end of word.") }

   navigator.Pop()
   if !navigator.TryPush('d') { t.Errorf("push should work") }
   if !navigator.EndOfWord() { t.Errorf("ad is end of word.") }
}
