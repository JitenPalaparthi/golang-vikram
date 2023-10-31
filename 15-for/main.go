package main

func main() {
	// normal loop
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		print(i, " ")
	// 	}
	// }

	//0 1 1 2 3 5 8 13 21 34

	a, b := 0, 1
	c := b
	//num := 100
	print(a, " ", b)
	for i := 1; i <= 10; i++ {
		c = b
		b = a + b
		print(" ", b)
		a = c
	}

}

// for
// for range
