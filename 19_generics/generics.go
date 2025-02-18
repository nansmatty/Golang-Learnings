package main

import "fmt"

// Passing any is not preferable and also you can pass the "interface{}" in the placec of any the meaning is same

// comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types). The comparable interface may only be used as a type parameter constraint, not as the type of a variable.

func printSlice[T comparable, N string](items []T, name N) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type stack[T int | string | bool] struct {
	elements []T
}

func main() {
	fmt.Println("Generics")
	nums := []int{1, 2, 3}
	// names := []string{"golang", "typescript"}
	bools := []bool{true, false}

	// printSlice(names)
	// printSlice(nums)
	printSlice(bools, "JD")

	myStack := stack[int]{
		elements: nums,
	}

	fmt.Println(myStack)

}
