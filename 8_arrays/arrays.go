package main

import "fmt"

func main() {

	//Use Cases of Array

	// Know the fixed data size, that is predictable
	// Memory optimization
	// Constant time access

	// Numbered sequence of specific length
	// When create an array with specific length if fills with zeros (In case the array type is integer)
	var nums [4]int

	nums[0] = 1

	fmt.Println(nums)

	//In case of array of strings with specific length fills with empty space

	var charVal [5]string

	charVal[0] = "narayan"

	fmt.Println(charVal)

	//In case of array of booleans with specific length fills with false value

	var vals [5]bool

	fmt.Println(vals)

	//array length
	// fmt.Println(len(nums))

	// Direct insert values while decalring the variable
	nums1 := [4]int{1, 2, 3, 4}

	fmt.Println(nums1)

	//2D arrays

	nums2d := [2][2]int{{1, 2}, {3, 5}}

	fmt.Println(nums2d)

}
