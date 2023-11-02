package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr1 [20]int
	fmt.Println(arr1)
	// assign random numbers to an array
	for i := 0; i <= len(arr1)-1; i++ {
		arr1[i] = rand.Intn(100) // it returns a randon number
	}
	fmt.Println(arr1)
	sum := 0

	for i := 0; i <= len(arr1)-1; i++ {
		sum += arr1[i]
	}

	fmt.Println(sum)
	// for array range loop gives two return values
	// index and value

	for i, v := range arr1 {
		fmt.Println("Index:", i, "Value:", v)
	}

}

// for range loop
// arrays, slices, maps, channels
