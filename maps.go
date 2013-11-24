package main

import "fmt"

func main() {
  a := []string{"1", "2"}  
  fmt.Println(a)

  // make(map[key_type]value_type)
  m := map[int]string{1: "one", 2: "two"}
  fmt.Println(m)
}
