package main

import (
	"fmt"
)

func main() {
	// int to float
	i := 123               // int
	iFloated := float64(i) //float64
	fmt.Println(i, iFloated)

	// float to int
	f := 122.07
	iIntegered := int(f)
	fmt.Println(f, iIntegered)

	// interface to string
	iface := map[string]interface{}{"hello": "world"}
	hello := iface["hello"].(string)
	fmt.Println(iface["hello"], hello)
}
