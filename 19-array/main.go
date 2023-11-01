// Process some memory and a thread
// Code Segment  --> data that is for the application itself(bin data)
// Data Segment  --> generally global variables
// Stack memory  -->
// Heap memory

// the below stuff is little different from programming language to language. But above 95 % they are same across programming languages

package main

import "fmt"

var n int // global are stored in Data segment (most of the cases)

func main() {
	// var num1 int = 100 // local variable

	// var num2 int = 200

	// {
	// 	var num3 = 400
	// 	{
	// 		var float1 float32 = 500.
	// 	} // end of the scope

	// }

	// num5 := 600
	// var str string = "Hello Golang"
	// const pi = 3.14                           // are stored in Code segment (most of the cases)
	// str = "Hello world! I am learning Golang" // mutate the variable str
	// num = 123123                              // mutate the variable num

	var arr1 [3]int // length of an array is 3. The index starts from 0

	arr1[0] = 100
	arr1[1] = 101
	arr1[2] = 102
	//arr1[3] = 103 // cannot access 4th element as the length of the array is 3

	var arr2 [3]int // type inference work here.
	fmt.Println(arr1)
	fmt.Println(arr2) // prints [0 0 0] becase type inference works here. arr2 is not assigned

	arr3 := [5]int{12, 13, 14, 15, 16} // shorthand declaration
	fmt.Println(arr3)

	arr4 := [...]int{10, 15, 20, 25} // compiler evaluates the size of this array based on values
	fmt.Println(arr4)
	// iterating through an array

	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

} // the scope of main ends

// array is collection of elements of same type
// index of an array starts from zero
// array is fixed size.the length of an array is fixed. That means you cannot inrease the size of an array , once it is declared
