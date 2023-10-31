package main

import (
	"fmt"
)

func main() {

	// var age uint8 = 10

	// var num uint64 = uint64(age) // type casting

	// println(num)
	// //var num1 int = int(any1)
	// // any type variable cannot be type casted
	// // instead ,can do type assertion

	// var any1 any
	// fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))

	// any1 = age
	// fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))

	// // var age2 int8 = any1.(uint8)
	// // println(age2)
	// num = uint64(any1.(uint8))

	// any1 = true
	// fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))

	// any1 = "Hello"
	// fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))

	// var str1 string = any1.(string)

	// fmt.Println(str1)

	var any2 any
	any2 = 123.45
	switch i := any2.(type) { // switch type switch
	case nil:
		fmt.Println(i, "is nil")
	case int:
		var i1 int
		i1 = i
		fmt.Println("Double the value is:", i1*2)
	case float64:
		var i1 float64
		i1 = i
		fmt.Println("Double the value is:", i1*2)
	case bool:
		var b1 bool
		b1 = i
		fmt.Println(b1)
	default:
		fmt.Println("invalid type")
	}
}

// asserting
// any.(t)
// 1.18
// boxing and unboxing
