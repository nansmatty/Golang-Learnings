package main

import "fmt"

const age = 30

func main() {
	const name string = "Golang"

	// name = "Javascript"

	//Constant grouping
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(port)
	fmt.Println(host)
}
