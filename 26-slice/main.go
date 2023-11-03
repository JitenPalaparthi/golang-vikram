package main

import "fmt"

func main() {
	var slice1 []int // no length is given as it is a slice declaration
	//var num1 int     // it has declaration and instantiation becase int is a defined, primitive type.
	// Allocating required memory is instantiation,
	// usually in object oriented programming languages, new is used to instantiate

	if slice1 == nil {
		println("slice is nil; that means it has not instantiated")
	} else {
		fmt.Println(slice1)
	}

	// how to instantiate slice

	// 1. at the time of declartion
	slice2 := []int{10, 11, 12} // declaration and instantiation as a shorthand stuff
	// arr2 := [3]int{10, 11, 12}
	// arr3 := [...]int{10, 11, 12}

	if slice2 == nil {
		println("slice is nil; that means it has not instantiated")
	} else {
		fmt.Println(slice2)
	}

	// var slice3 []int // declaration
	// slice3 = make([]int, 5) // this is instantiation

	slice3 := make([]int, 5) // this is declaration and instantiation
	fmt.Println("Address of slice:", &slice3[0])
	// fmt.Println("Address of slice:", &slice3[1])
	// fmt.Println("Address of slice:", &slice3[2])

	fmt.Println("Values slice3:", slice3)          // [0 0 0 0 0]
	fmt.Println("Length slice3:", len(slice3))     // 5
	fmt.Println("Capacity of slice3", cap(slice3)) // 5
	// if you dont mention the capacity explicitly the length and cap are same [ There is another clause]

	slice3[0] = 100
	slice3[1] = 101
	slice3[2] = 102
	slice3[3] = 103
	slice3[4] = 104
	slice3 = append(slice3, 105) // grow a slice
	fmt.Println("Address of slice3:", &slice3[0])
	fmt.Println("Values slice3:", slice3) // [0 0 0 0 0]
	fmt.Println("Cap slice3:", cap(slice3))
	slice3 = append(slice3, 106)          // grow a slice
	slice3 = append(slice3, 107)          // grow a slice
	slice3 = append(slice3, 108)          // grow a slice
	slice3 = append(slice3, 109)          // grow a slice
	fmt.Println("Values slice3:", slice3) // [0 0 0 0 0]
	fmt.Println("Address of slice3:", &slice3[0])
	slice3 = append(slice3, 110)

	fmt.Println("Cap slice3:", cap(slice3))
	fmt.Println("Values slice3:", slice3) // [0 0 0 0 0]
	fmt.Println("Address of slice3:", &slice3[0])

	//slice3[5] = 105

	// arr1 := [1]int{100}
	// arr1[0] = 101
	// arr1[1] = 200
	sum := 0
	for i, v := range slice3 {
		sum += v // sum = sum+v
		fmt.Println("Index:", i, "Value:", v)
	}
	fmt.Println("Sum of elements in slice3:", sum)

	var i = 100
	var j = i // deep copy

	fmt.Println(i, j)
	j = 200

	fmt.Println(i, j)

	slice4 := slice3 // shallow copy
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)
	slice4[0] = 10000
	fmt.Println("slice3:", slice3)

}

// slice can grow or shrink at runtime
// declaration and instantiation
// nil is used to check whether a variable is instantiated or not
// nil can used for slice, map, pointer, chan, interface/any
// make is a built in function to instantiate only slice, map, chan
// slice header
/*
	Ptr: The pointer to the actual memory
	Len: Number of elements in the slice
	Cap: Number of elements that the slice can accomidate
*/
