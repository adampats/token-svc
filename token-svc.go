package main

import (
  "fmt"
  "log"
  "net/http"
  "math/rand"
  "time"
)

var port string = ":8080"
var tokenlen = 40
var r *rand.Rand

func gimmeToken(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, generateToken(tokenlen))
}

func generateToken(length int) string {
  const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  result := make([]byte, length)
  for i := range result {
  	result[i] = chars[r.Intn(len(chars))]
  }
  return string(result)
}

func main() {
  r = rand.New(rand.NewSource(time.Now().UnixNano()))
  http.HandleFunc("/", gimmeToken)
  err := http.ListenAndServe(port, nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
