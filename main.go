package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
}

func main() {
	p1 := person{
		First: "Jed",
		Last:  "Zeins",
	}
	p2 := person{
		First: "Claude",
		Last:  "Zeins",
	}

	persons := []person{p1, p2}

	bs, err := json.Marshal(persons)
	// turns to json string?////

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(bs))

	persons2 := []person{}

	// needs pointer of object
	err = json.Unmarshal(bs, &persons2)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Back into a go data structure", persons2)

}
