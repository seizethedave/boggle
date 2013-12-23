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
