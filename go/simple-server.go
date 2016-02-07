package main

import (
  "net/http"
  "fmt"
  "time"
  "math/rand"
)

var hiscore = []int {900, 800, 700, 600, 500,
  400, 300, 200, 100, 10}

func main() {
  http.HandleFunc("/", topHandler)
  http.HandleFunc("/list", rankingHandler)
  http.HandleFunc("/hiscore", hiscoreHandler)
  http.ListenAndServe(":8080", nil)
}

func topHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "simple hiscore ranking server\n")
  fmt.Fprintf(w, "%s", time.Now())
}

func  dispHiscore(w http.ResponseWriter) {
  fmt.Fprintf(w, "Ranking is ...\n")
  for i, v := range hiscore {
    fmt.Fprintf(w, "%d %d\n", i, v)
  }
  fmt.Fprintf(w, "%s\n", time.Now())
}

func rankingHandler(w http.ResponseWriter, r *http.Request) {
  dispHiscore(w)
}

func hiscoreHandler(w http.ResponseWriter, r *http.Request) {
// 本来 hiscore は client -> server 伝送されるものだが、
// 手を抜いて、ここでは乱数で代用する。
  var hs int = rand.Intn(1000)
  fmt.Fprintf(w, "hiscore candidate is %d\n", hs)

  for i, v := range hiscore {
    if hs >= v {
      // new hiscore
      for j := len(hiscore)-1; j>=i+1; j-- {
          hiscore[j] = hiscore[j-1]
      }
      hiscore[i] = hs
      break;
    }
  }

  dispHiscore(w)
}
