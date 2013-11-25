package main

import "fmt"

//funcs
func plus(a int, b int) int {
  return a + b
}

//multiple return values
func vals() (int, int) {
  return 88, 77
}

//Varidatic function
func sum(nums ...int) int {
  total := 0
  for i := 0; i < len(nums); i++ {
    total = total + nums[i]
  }

  return total
}

func main() {
  result := plus(9,3)
  fmt.Println(result)

  a, b := vals()

  fmt.Println(a)
  fmt.Println(b)

  total := sum(vals())
  fmt.Println(total)
}
