package main

import (
	"fmt"
	"reflect"
)

func main() {

	var arr1 [3]int // [3]int [0 0 0]
	arr2 := [4]int{10, 11, 12, 13}
	arr3 := [...]int{0, 0, 0, 0, 0}

	//arr1 = arr2 // not ok; because both of them are two different types
	for i := range arr1 {
		arr1[i] = arr2[i]
	}
	print("{")
	//var i, v int
	for _, v := range arr1 { // _ blank identifier
		print(v, " ")
	}

	print("}")

	// for i, _ := range arr1 { // _ blank identifier
	// 	print(i, " ")
	// }

	println()

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	str := fmt.Sprint(arr1) // This converts the whole array to string
	fmt.Println(str, " Type of arr1:", reflect.TypeOf(str))

}

// The type system is Go is very strict
// Variables of different types cannot be implicitely conveted
// Type casting also cannot be done , but there is a way to do it (will discuss in the next concept)

// difference between fmt.Println and println
// println is a builtin function
// it is easy to print. But it cannot print(from the developers perspective)
// fmt.Println : formatted print, to help engineer , to see data
// println is a very good performer (from allocations perspective)
// fmt.Println usually allocates on heap; Heap allocations are less performers than stack
