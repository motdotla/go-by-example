package main

import "fmt"

func main() {
  i := 1
  for i<=3 {
    fmt.Println(i)
    i = i + 1
  } 

  for j := 5; j <= 12; j++ {
    fmt.Println(j)
  }

  for {
    fmt.Println("dude")
    if i >= 100 {
      break
    }
    i = i + 1
  }
}
