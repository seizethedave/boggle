package boggle

import "testing"

func TestOneWordNode(t *testing.T) {
   node := NewSuffixTree("ape")

   if node.children['a'].character != 'a' {
      t.Errorf("Expected character to be a and not '%c'", node.character)
   }

   if node.endOfWord {
      t.Errorf("First node should not be end of word.")
   }

   if node.children == nil || len(node.children) == 0 {
      t.Errorf("Expected children to be non-nil and non-empty.")
   }

   if !node.children['a'].children['p'].children['e'].endOfWord {
      t.Errorf("Last node should be end of word.")
   }
}

func TestAFewWords(t *testing.T) {
   root := NewSuffixTree("ape")
   root.Add("apply")
   root.Add("apes")

   apeNode := root.children['a'].children['p'].children['e']
   if !apeNode.endOfWord {
      t.Errorf("Expected endOfWord at end of ape.")
   }

   apesNode := root.children['a'].children['p'].children['e'].children['s']
   if !apesNode.endOfWord {
      t.Errorf("Expected endOfWord at end of apes.")
   }

   applyNode :=
    root.children['a'].children['p'].children['p'].children['l'].children['y']
   if !applyNode.endOfWord {
      t.Errorf("Expected endOfWord at end of apply.")
   }
}
