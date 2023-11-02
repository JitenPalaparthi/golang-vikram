package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr1d := [2]int{1, 2}
	arr2d := [3][3]int{{100, 200, 300}, {300, 400, 500}, {600, 700, 800}}
	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	arr4d := [2][2][2][2]int{{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}, {{{9, 10}, {11, 12}}, {{13, 14}, {15, 16}}}}
	fmt.Println(arr1d)
	fmt.Println(arr2d)
	fmt.Println(arr3d)
	fmt.Println(arr4d)

	fmt.Println(arr2d[1][1])           // 1,2
	fmt.Println(reflect.TypeOf(arr2d)) // [2][2]int
	fmt.Println("Focus here")
	for _, v1 := range arr2d {
		//fmt.Println("Type of v:", reflect.TypeOf(v), "Value:", v)
		sum := 0 // a new variable is created for every outer iteration
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
			sum += v2
		}
		fmt.Println(sum)
	}
	// [2]int
	//
}

// In a multidimensional array , each dimension require one loop..
// if it is 3d array , to iterate and access each element , it reqire three loops

// array transpose of an array
