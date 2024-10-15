package main

import (
	"fmt"
)

func main() {
	fmt.Println("Letsssssssss gooooooooo")

	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println(i)
	// case 2:
	// 	fmt.Println(i)
	// case 3:
	// 	fmt.Println(i)
	// case 4:
	// 	fmt.Println(i)
	// default:
	// 	fmt.Println("Other")
	// }

	// Multiple condition switch statement

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It is weekend!!! Aeee Daaru kidhar hai kidhar hai daaaruuuu")
	// default:
	// 	fmt.Println("It is workday! Boring.........")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It is an integer")
		case string:
			fmt.Println("It is a string")
		case bool:
			fmt.Println("It is a boolean")
		default:
			fmt.Println("Other", t)

		}
	}

	whoAmI(79)

}
