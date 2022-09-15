package main

import (
	"fmt"
	"log"

	"github.com/ianlopshire/go-fixedwidth"
)

func main() {
	// define some data to encode
	people := []struct {
		ID        int     `fixed:"1,5"`
		FirstName string  `fixed:"6,15"`
		LastName  string  `fixed:"16,25"`
		Grade     float64 `fixed:"26,30"`
		Age       uint    `fixed:"31,33"`
		Alive     bool    `fixed:"34,39"`
	}{
		{1, "Ian", "Lopshire", 99.5, 20, true},
	}

	data, err := fixedwidth.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
	// Output:
	// 1    Ian       Lopshire  99.5020 true
}