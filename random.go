package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // https://groups.google.com/forum/#!topic/golang-nuts/TzsEceDLGI0
	fmt.Println("Random number: ", rand.Intn(100))
}
