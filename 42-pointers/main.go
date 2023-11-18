package main

import "fmt"

func main() {
	num1 := 100                                         // declaring and assigning a value
	fmt.Println("num1 before calling incr func:", num1) // prints
	incr(num1)                                          //
	fmt.Println("num1 after calling incr func:", num1)  // does it print 101?
	ptr1 := &num1                                       // shorthad declaration of pointer
	incrPtr(ptr1)                                       // call by reference
	incrPtr(&num1)                                      // call by reference
	fmt.Println("num1 after calling incr func:", num1)  // does it print 101?

	// var x int = 200
	// var y int = 300
	// var ptr1 *int = &x // address to ptr1
	// fmt.Println(ptr1)  // address of x
	// fmt.Println(*ptr1) // Value at the address
	// ptr1 = &y          // now ptr holds the address of y
	// fmt.Println(ptr1)
	// fmt.Println(*ptr1)

}

// by default func calls takes arguments which are copies of variables. Pass or call by value
func incr(x int) {
	x = x + 1
	//fmt.Println(x)
}

func incrPtr(x *int) { // call by reference
	*x = *x + 1
	//fmt.Println(x)
}

// What are parameters
// what are return types/parameters
// how does func crate variables
// each fun creates its own set of variables like parameters , return types, local variables etc..
// usually their scope is with in the func's scope. Pointers or other reference types are different
