package main

import "fmt"

func main() {
	// unconditional loop
	// for {
	// 	println("Hello")
	// }

	i := 1 // init

	for i <= 10 { // condition
		if i%2 == 0 {
			println(i)
		}
		i++
	}

	for i := 1; ; {
		if i > 10 {
			break
		}
		fmt.Print(i, " ")
		i++
	}

	fmt.Println("another type")

	for i := 1; ; i++ {
		if i > 10 {
			break
		}
		fmt.Print(i, " ")
		//i++
	}

	fmt.Println("another variation")

	i = 1
	for {
		if i > 10 {
			break
		}
		println(i)
		i++
	}

	fmt.Println("continue")

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // it skips the current iteration and continue with the next iteration
		}
		print(i, " ")
	}

}
