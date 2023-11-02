package main

import "fmt"

func main() {
	greet()
	greet()
	greetP("Hello Golang programming!")
	greetP("Hello Golang.Go is a better performer than many other programming languages")
	add(10, 20)
	c := addP(10, 20)
	fmt.Println(c)
	fmt.Println("Square of c:", c*c)
	sum1, sub1 := calc(10, 20)
	fmt.Println("Sum:", sum1, "Substract:", sub1)
	_, sub2 := calc(10, 20)
	fmt.Println("Substract:", sub2)
	sum3, _ := calc(10, 20)
	fmt.Println("Sum:", sum3)

	c1, c2, c3 := calc3(10, 20)
	fmt.Println("Sum:", c1, "Sub:", c2, "Mul:", c3)

	c1, _, c3 = calc3(200, -200)
	fmt.Println("Sum:", c1, "Mul:", c3)

}

func greet() {
	fmt.Println("Hello Golang!")
}

func greetP(msg string) {
	fmt.Println(msg)
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func addP(a int, b int) int {
	c := a + b
	//fmt.Println(c)
	return c
}

func calc(a, b int) (int, int) {
	c1 := a + b
	c2 := a - b
	//return a + b, a - b
	return c1, c2
}

func calc3(a, b int) (sum int, sub int, mul int) {
	sum = a + b
	sub = a - b
	mul = a * b
	//return a + b, a - b
	return sum, sub, mul
}

// What: func is a block of execution
// Why: if you want a block of code to be called multiple times(0-n times), you can write it as a func and call it
// How: each functon starts with func keyword.
// 		there can be 0-n input parameters and 0-n return values
// there should be a caller for a function to execute
// funcs may or may not return values
// funcs in go can return multiple values . No other programming language(as far as I see) has this feature that returns multiple return values.Usually tuples are used for other programming languages
