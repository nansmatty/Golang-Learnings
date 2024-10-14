package main

import "fmt"

func main() {
	fmt.Println("Letsssssssss gooooooooo")

	// age := 16

	// if age >= 18 {
	// 	fmt.Println("The person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("This person is teenager")
	// } else {
	// 	fmt.Println("This person is a kid")
	// }

	// logicaal operators

	// Go doesn't have the ternary operator.........................

	role := "admin"
	hasPermission := true

	if role == "admin" && hasPermission {
		fmt.Println("Yes")
	}

	if role == "admin" || hasPermission {
		fmt.Println("Yes")
	}

	// variable declartion in if statement

	if age := 15; age >= 18 {
		fmt.Println("Person is an adult: ", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager: ", age)
	} else {
		fmt.Println("Person is a kidL ", age)
	}

}
