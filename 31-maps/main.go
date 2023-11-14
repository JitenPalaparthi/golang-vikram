package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// map1 := make(map[string]any)                                           // declaration and instantiation. Shorthand declaration
	// map2 := map[string]int{"India": 91, "Pakisthan": 92, "Bangladesh": 93} // shorthand declaration
	// fmt.Println(map1, map2)
	// var map3 map[string]any // it is a nil map. It is declared but not instantiated
	// if map3 == nil {
	// 	fmt.Println("nil map3")
	// } else {
	// 	fmt.Println(map3)
	// }

	slice1 := make([]int, 100)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(50)
	}
	fmt.Println(slice1)

	fmt.Println("----Find duplicate values from a slice")
	DuplicatePrint(slice1)
	fmt.Println("---------------------------------------")
	fmt.Println("----Find duplicate values from a slice using map")
	DuplicatesPrintUsingMap(slice1)

	// map1 := make(map[string]int)
	// map1["India"] = 91
	// map1["Pakithan"] = 92

	// v, ok := map1["India"] // it can return two
	// //if ok == false {
	// if ok {
	// 	fmt.Println("key India existed, so the value is", v)
	// } else {
	// 	fmt.Println("Key does not exist")
	// }

}

func DuplicatePrint(slice1 []int) {
	duplicate := make([]int, 0)
	loopCounter := 0
	for i := 0; i < len(slice1); i++ {
		loopCounter++
		value := slice1[i]
		count := 1
		isDuplicate := false
		for k := 0; k < len(duplicate); k++ {
			loopCounter++
			if value == duplicate[k] {
				isDuplicate = true
			}
		}
		if !isDuplicate {
			duplicate = append(duplicate, value)
		} else {
			continue
		}
		for j := i + 1; j < len(slice1); j++ {
			loopCounter++
			if value == slice1[j] {
				count++
			}
		}
		fmt.Println("Value:", value, "Count:", count)
	}
	fmt.Println("Total Iterations using normal loops:", loopCounter)
}

func DuplicatesPrintUsingMap(slice1 []int) {
	totalIterations := 0
	map1 := make(map[int]int)
	for i := 0; i < len(slice1); i++ {
		totalIterations++
		v, ok := map1[slice1[i]]
		if !ok {
			map1[slice1[i]] = 1
		} else {
			map1[slice1[i]] = v + 1
		}
	}
	for k, v := range map1 {
		fmt.Println("Value:", k, "Count:", v)
	}
	fmt.Println("Total Iterations using normal map:", totalIterations)
}

// mostly to avoid multiple loops
// to store data in memory
// to map values to the keys. As and where unique keys to be maintained.

// Maps are not for when there are duplicate keys to be maintained
// When order of elements is important. There is another orderd maps but it is from a thirdpary package. It is not as efficient as builtin maps
