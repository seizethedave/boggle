package boggle

import (
   "io"
   "bufio"
   "log"
)

type BoggleDictionary struct {
   root *SuffixTreeNode
}

func NewBoggleDictionaryWithWords(words ...string) (dict *BoggleDictionary) {
   dict = &BoggleDictionary { }
   dict.root = &SuffixTreeNode { }
   for _, word := range words {
      dict.root.Add(word)
   }
   return
}

func NewBoggleDictionaryFromReader(reader io.Reader) (dict *BoggleDictionary) {
   dict = &BoggleDictionary { }
   dict.root = &SuffixTreeNode { }

   scanner := bufio.NewScanner(reader)

   for scanner.Scan() {
      dict.root.Add(scanner.Text())
   }

   if err := scanner.Err(); err != nil {
      log.Fatal(err)
   }

   return
}

func (dict *BoggleDictionary) Contains(word string) bool {
   node := dict.root.Seek(word)
   return node != nil && node.endOfWord
}
