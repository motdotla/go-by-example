package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func main() {
	var person Person

	person.Name = "Scott Motte"
	person.Sex = "male"

	marshaled_person, _ := json.Marshal(person)

	log.Println(string(marshaled_person))
}
