package main

import (
  "fmt"
  "math"
)

const s string = "constant"

func main() {
  fmt.Println(s)

  const n float64 = 500000000

  fmt.Println(math.Sin(n))
}
