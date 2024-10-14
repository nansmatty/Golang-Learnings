package main

// for -> only construct in go for looping

func main() {

	// while loop fashion
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//infinite loop
	// for {
	// 	println("1")
	// }

	// classic for loop
	// for i := 0; i < 10; i++ {
	// println(i)

	// break
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	println(i)
	// }

	// 1.22 version in golang feature ==== Range

	for i := range 3 {
		println(i)
	}

}
