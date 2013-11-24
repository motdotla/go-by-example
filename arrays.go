package main

import "fmt"

func main() {
  var array [5]int
  fmt.Println(array)

  array[4] = 44
  fmt.Println(array)
  fmt.Println(array[4])


  b := [5]int{1,2,3,4,5}
  fmt.Println(b)

  twoD := [2][3]int{{1,2,3}, {4,5,6}}
  fmt.Println(twoD)
}
