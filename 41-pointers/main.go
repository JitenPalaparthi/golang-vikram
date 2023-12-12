package main

import "fmt"

func main() {
	var num1 int = 100    // normal variables hold values
	var ptr1 *int = &num1 // pointer holds the address of an another variable

	fmt.Println("Address of num1:", &num1)
	fmt.Println("Address of num1:", ptr1) // the address of num1
	fmt.Println("Value at that address:", *&num1)
	fmt.Println("Value at the address:", *ptr1) // * denotes value at the address
	fmt.Println("Ptr has its own address:", &ptr1)

	// pointer dereference
	*ptr1 = 200
	fmt.Println("Value of num1:", num1)

	// type inference

	var ptr2 *int // ptr2 is a pointer but it does not hold any address..

	// what is the value of ptr2
	fmt.Println("ptr2:", ptr2)
	fmt.Println("Value at ptr2:", *ptr2)

	// var ptr3 *int
	// *ptr3 = *ptr3 + 1
	// fmt.Println("ptr3:", *ptr3)

}

// If there is a pointer , it must contain the address of a variable
// * says the value at the address
// *ptr is dereference
// if a pointer does not hold any address , then it contains nil
