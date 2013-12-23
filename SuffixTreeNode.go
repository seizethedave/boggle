package boggle

type SuffixTreeNode struct {
   character rune
   endOfWord bool
   children map[rune]*SuffixTreeNode
}

func NewSuffixTree(word string) (node *SuffixTreeNode) {
   prefix, suffix := rune(word[0]), word[1:]
   node = &SuffixTreeNode { character: prefix }
   node.Add(suffix)
   return
}

func (node *SuffixTreeNode) Add(word string) {
   if 0 == len(word) {
      // We've recursed into an empty suffix. The word ends here.
      node.endOfWord = true
      return
   }

   prefixRune, suffix := rune(word[0]), word[1:]
   var prefixNode *SuffixTreeNode

   if nil == node.children {
      // We avoid actually creating the children map util this point so leaf
      // nodes don't eat unnecessary memory.

      prefixNode = &SuffixTreeNode { character: prefixRune }
      node.children = make(map[rune]*SuffixTreeNode)
      node.children[prefixRune] = prefixNode
   } else {
      prefixNode = node.children[prefixRune]

      if prefixNode == nil {
         // The children map has been created, but doesn't contain this rune.
         prefixNode = &SuffixTreeNode { character: prefixRune }
         node.children[prefixRune] = prefixNode
      }
   }

   // Call into the new node's Add method with the remainder of the word.
   // (A ten-character word will result in ten nested SuffixTreeNodes.)

   prefixNode.Add(suffix)
}

