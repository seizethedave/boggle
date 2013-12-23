package boggle

type DictionaryNavigator interface {
   TryPush(rune) bool
   Pop()
   EndOfWord() bool
}

type BoggleDictionaryNavigator struct {
   node *SuffixTreeNode
   history []*SuffixTreeNode
}

func NewBoggleDictionaryNavigator (rootNode *SuffixTreeNode) *BoggleDictionaryNavigator {
   return &BoggleDictionaryNavigator { node: rootNode }
}

func (navigator *BoggleDictionaryNavigator) TryPush(character rune) bool {
   nextNode := navigator.node.children[character]

   if nextNode != nil {
      navigator.history = append(navigator.history, nextNode)
      navigator.node = nextNode
      return true
   }

   return false
}

func (navigator *BoggleDictionaryNavigator) Pop() {
   if len(navigator.history) == 0 {
      panic("Too many pops.")
   }

   navigator.history = navigator.history[:len(navigator.history) - 1]
}

func (navigator *BoggleDictionaryNavigator) EndOfWord() bool {
   return navigator.node.endOfWord
}

// Implements a fake navigator which basically says "sure" with any question.
type DumbDictionaryNavigator struct { }

func (navigator *DumbDictionaryNavigator) TryPush(character rune) bool {
   return true
}

func (navigator *DumbDictionaryNavigator) Pop() { }

func (navigator *DumbDictionaryNavigator) EndOfWord() bool {
   return true
}
