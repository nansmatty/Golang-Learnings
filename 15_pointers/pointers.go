package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("In change num :", *num)
}

func main() {
	fmt.Println("Pointers")

	num := 3

	changeNum(&num)

	fmt.Println("After change num in main : ", num)
	// fmt.Println("Memory address of num variable : ", &num)

}
