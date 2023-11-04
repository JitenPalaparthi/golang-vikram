package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr1 := [10]int{}
	for i := range arr1 {
		arr1[i] = rand.Intn(100)
	}

	//sum1, sub1, mul1 := calc(arr1) // by default a slice to be passes as an input argument but not array
	//[]int is different [10]int
	slice1 := arr1[:] // this does or converts an array to a slice. This is a memory copy of array to slice , shallow copy
	sum1, sub1, mul1 := calc(slice1)
	fmt.Println(sum1, sub1, mul1)

	sum2, sub2, mul2 := calc(arr1[:])
	fmt.Println(sum2, sub2, mul2)

	arr2 := [...]int{10, 11, 12} // array declaration and assign values
	arr3 := arr2                 // deep copy
	arr3[2] = 112                // it does not change arr2

	fmt.Println(arr2)
	fmt.Println(arr3)

	slice2 := arr2[:] // shallow copy , the ptr if arr2 , len of arr2 and cap of arr2 is copied to slice2 header
	slice2[2] = 11112

	arr2[0] = 11111111

	fmt.Println(arr2)
	fmt.Println(slice2)
	slice2 = append(slice2, 4444) // it reallocates but the newly allocated memory is divergent to the arr2
	slice2[0] = 0
	fmt.Println("arr2", arr2)
	fmt.Println("slce2", slice2)

}

// calc func takes a slice as a input argument and
// it returns three integers.1st one is sum of slice. 2nd one is substraction of slice
// the 3rd one is multiplication is slice.
func calc(s []int) (sum int, sub int, mul int) {
	//sum, sub, mul := 0, 0, 1 // creation and assign values
	if mul == 0 { // no point is checking mul is 0 or not(by default num is 0) but to explain scope.
		mul = 1
	}
	for _, v := range s { // v is created inside for-range scope.
		sum += v // add = add+v // mutate
		sub -= v // mutate
		mul *= v // mutate
	} // v cannot be used beyond this point
	// fmt.Println(v)
	return sum, sub, mul
}
