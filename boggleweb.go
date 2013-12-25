package boggleweb

import (
   "net/http"
   "html/template"
   "encoding/json"
   "os"
   "bufio"
   "log"

   "boggle"
)

var indexTemplate = template.Must(template.ParseFiles("base.html", "index.html"))
var boggleDictionary *boggle.BoggleDictionary

func init() {
   http.HandleFunc("/api/boggle/words", solve)
   http.HandleFunc("/", home)

   dictFile, err := os.Open("dict.txt")

   if err != nil {
      panic(err)
   }

   defer dictFile.Close()
   reader := bufio.NewReader(dictFile)
   boggleDictionary = boggle.NewBoggleDictionaryFromReader(reader)
}

func home(w http.ResponseWriter, r *http.Request) {
   if err := indexTemplate.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
   }
}

func solve(w http.ResponseWriter, r *http.Request) {
   board := boggle.NewBoardFromString(r.FormValue(""))

   log.Printf("Dict: %q", boggleDictionary)
   words := board.ScanAll(boggleDictionary)

   if result, err := json.Marshal(words); err == nil {
      w.Header().Set("Content-Type", "application/json")
      w.Write(result)
   } else {
      http.Error(w, err.Error(), http.StatusInternalServerError)
   }
}
