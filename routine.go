package main

import (
  "fmt"
  "time"
  "math/rand"
)

func do(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}

func main() {
  for i := 0; i < 10; i++ {
    go do(0)
  }
  var input string
  fmt.Scanln(&input)
}
