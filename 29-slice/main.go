package main

import "fmt"

func main() {

	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	arr1 := [3]int{}
	arr1 = [3]int(slice1) // type casting of slice to array. This is not a shallow copy
	slice2 := arr1[:]     // converting arr to slice
	fmt.Println("slice1:", slice1)
	fmt.Println("arr1:", arr1)
	fmt.Println("slice2:", slice2)

	// copy a part of slice

	slice3 := slice1[3:5] // it is a shallow copy of 4th and 5th elements

	fmt.Println("slice3:", slice3)
	fmt.Println("len slice3:", len(slice3))
	fmt.Println("cap slice3:", cap(slice3))
	slice3[0] = 99999999
	fmt.Println("slice3:", slice3)
	fmt.Println("slice1:", slice1)

	slice4 := slice1[4:] // 4th element to the end
	slice5 := slice1[:4] // from 0th element to the 3rd element
	fmt.Println("slice4:", slice4)
	fmt.Println("slice5:", slice5)

	slice6 := []string{"Bangalore", "Hyderabad", "Chennai", "Trivandrum"}
	fmt.Println("slice6:", slice6)

	// deep copy with a slice

	slice7 := make([]int, len(slice1))
	copy(slice7, slice1) // all elements of slice1 are copied to slice7. This is a deep copy
	// any change on slice7 or vice versa does not impact any of the slices
	// slice1 and slice7 are two different copies
	fmt.Println("slice7:", slice7)
}

// you can create a slice from an array but not the other way directly
// how to create an array from a slice is direct assignment of each element or type caste
