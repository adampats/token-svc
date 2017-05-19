package main

import (
  "os"
  "fmt"
  "log"
  "net/http"
  "math/rand"
  "time"
  "strconv"
)

var port = 8080
var tokenlen = 40
var r *rand.Rand

func printToken(w http.ResponseWriter, r *http.Request) {
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
  // assign listening port number to 1st command line argument, if provided
  args := os.Args[1:]
  if len(args) != 0 {
    port, _ = strconv.Atoi(args[0])
  }
  portSuffix := fmt.Sprintf(":%v", strconv.Itoa(port))

  r = rand.New(rand.NewSource(time.Now().UnixNano()))
  http.HandleFunc("/", printToken)
  fmt.Println("Listening on port " + portSuffix + "...")
  err := http.ListenAndServe(portSuffix, nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
