package main

import "fmt"

// range => iteration over data structures

func main() {
	fmt.Println("Ranges")

	//iteration over slices
	nums := []int{1, 3, 4, 6, 9, 8, 2}

	// iteration using for loop

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// }

	// iteration using range

	sum := 0
	for _, num := range nums {
		sum = sum + num
		// fmt.Println(index, " - ", num)
	}
	fmt.Println("Sum of Slice is : ", sum)

	// iteration over maps
	m := map[string]string{"fname": "John", "lname": "Doe"}

	for k, v := range m {
		fmt.Println(k, " - ", v)
	}

	// iteration oover string
	for index, char := range "golang" {
		fmt.Println(index, "--", string(char))
	}

}
