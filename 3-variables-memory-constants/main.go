package main

import (
	"fmt"
)

func main() {
	// local variables . local to main
	{ // ----------------------------------------------
		var (
			num1 int = 10
			num2 float32
			ok1  bool
			str1 string
		)
		fmt.Println(num1, num2, ok1, str1)
	} // ---------------------------------------------------
	fmt.Println(num1, num2, ok1, str1)

	fmt.Println(pi1, pi2, pi3, piConstant)

}

// global variables
var (
	num1       int = 100
	num2       float32
	ok1        bool
	str1       string
	piConstant float32 = 0.003
)

// constant

const (
	constantPi float32 = 0.003
	pi1        float32 = 3.14 + 0.003 // evaluation at compile time
	pi2        float32 = 3.14
	pi3        float32 = pi2 + constantPi
	// pi4        float32 = pi2 + piConstant // caanot be done as piConstant is a variable
)

// What is a constant?
// a value will never be changed
// is any expression has to evaluate the value of a constant 1.
// It has to be evaluated at compile time
// no evaluation happens at run time w.r.t. constants

// scope
// Where does local and global variables store ?
// A process has some memory
// memory is nothing but array of bytes

// code segment   ---> constants are stored here. Constants are stored in the binary
// data segment   ---> global variables are stored in data segment
// stack memory
// heap memory
