package main

import (
	"fmt"
	"maps"
)

// Maps => Objects in Javascript

func main() {

	// creating a map
	m := make(map[string]string)

	// Setting an element in map
	m["name"] = "golang"
	m["area"] = "backend"

	//Get an element
	fmt.Println(m)
	// fmt.Println(m["name"], m["area"])

	//IMP: if key does not exists in the map then it return empty in case string, zero in case of integer and false in case of boolean
	// fmt.Println(m["phone"])

	// delete a key value from map
	// delete(m, "area")

	// clear the entire map
	// clear(m)

	// fmt.Println(m)

	// second wat to create map
	m1 := map[string]int{"phone": 3483287432, "age": 27, "price": 20000000}
	m2 := map[string]int{"phone": 3483287432, "age": 27, "price": 20000000}

	// Equal check

	fmt.Println(maps.Equal(m1, m2))

	fmt.Println(m1)

	k, ok := m1["age"]

	fmt.Println(k)

	if ok {
		fmt.Println("All Ok")
	} else {
		fmt.Println("Not Ok")
	}

}
