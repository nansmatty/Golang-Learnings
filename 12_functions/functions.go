package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// Multiple value type returning statement showing
func getLanguages() (string, string, string) {
	return "golang", "javascript", "python"
}

// function which accepts another function

func processIt(fn func(a int) int) {
	fn(1)
}

// function which returns another function

func processIt2() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	fmt.Println("Functions")

	results := add(4, 9)

	fmt.Println(results)
	fmt.Println(getLanguages())

	// funtion which accepts another funtion

	fn := func(a int) int {
		return 4
	}

	processIt(fn)

	fn2 := processIt2()

	fn2(6)

}
