package main

import (
	"fmt"
	"slices"
)

//Slices are dynamic array
// most used construct

func main() {

	// uninitialized slice in nil
	// in js we use null in cse of golang it is nil
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	//second way to create a slice
	// var iniNums = make([]int, 0, 3)

	// Cap or capacity -> maximum numbers of elements can be fit

	//third way to create a slice....
	// iniNums := []int{}

	// iniNums = append(iniNums, 1)
	// iniNums = append(iniNums, 2)
	// iniNums = append(iniNums, 3)
	// iniNums = append(iniNums, 4)
	// iniNums = append(iniNums, 5)
	// iniNums = append(iniNums, 6)
	// fmt.Println(iniNums)
	// fmt.Println(cap(iniNums))

	//Copy function

	var numeric = make([]int, 0, 5)

	numeric = append(numeric, 5)

	var numeric2 = make([]int, len(numeric))

	copy(numeric2, numeric)

	fmt.Println(numeric, numeric2)

	// Slice Operator

	var nums = []int{64, 32, 56, 73}

	fmt.Println(nums[1:3])
	fmt.Println(nums[1:]) // upto end no variable will be omit
	fmt.Println(nums[:3]) // from start like if we

	// Slices package builtin standard golang library

	var nums1 = []int{1, 2, 3, 4}
	var nums2 = []int{1, 2, 3, 4}

	fmt.Println(slices.Equal(nums1, nums2))
	fmt.Println(slices.Equal(numeric, numeric2))

	// 2D Slices

	var nums2d = [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(nums2d)

}
