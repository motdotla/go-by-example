package main

import "fmt"

func main() {
  slice := []string{"dude", "ted", "bill"}
  fmt.Println(slice)

  slice[0] = "newdude"
  fmt.Println(slice)
}
