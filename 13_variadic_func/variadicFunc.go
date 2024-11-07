package main

import "fmt"

// Variadic function means a function which can accept n no. of variables

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {
	fmt.Println("Variadic Functions")

	// creating slice and passing through variadic func
	nums := []int{1, 3, 6, 3, 6, 3, 63, 5, 5, 7, 4, 3, 4, 4, 7, 44, 4, 3, 7, 44}

	result := sum(nums...)

	fmt.Println(result)

}
