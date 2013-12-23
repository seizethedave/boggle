package boggle

import "testing"

func TestOneWordNode(t *testing.T) {
   node := NewSuffixTree("ape")

   if node.character != 'a' {
      t.Errorf("Expected character to be a and not '%c'", node.character)
   }

   if node.endOfWord {
      t.Errorf("First node should not be end of word.")
   }

   if node.children == nil || len(node.children) == 0 {
      t.Errorf("Expected children to be non-nil and non-empty.")
   }

   if !node.children['p'].children['e'].endOfWord {
      t.Errorf("Last node should be end of word.")
   }
}
