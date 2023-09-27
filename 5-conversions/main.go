package main

import "fmt"

func main() {

	var num1 int8 = 100
	//var num2 int16 = num1 // no implicit conversion in go

	// assume java code , it is valid in jave
	// byte num1 = 100;
	// long num2 = num1;

	var num2 int16 = int16(num1) // type casting or explicit conversion

	fmt.Println(num1, num2)

	var num3 int16 = 23213
	// 00000000 00000000
	// 01011010 10101101

	var num4 uint8 = uint8(num3)
	// 10101101 --> if unit8 then 173 if int8 -

	fmt.Println(num3, num4)

	const pi = 3.14 // float64

	var vpi float32 = pi // float32

	// var ipi int32 = pi

	fmt.Println(vpi, pi)
}

// conversions or type casting
// converting one type to another type
// in general two types of conversions exist
// 1- implicit -> no implicit conversions in go except a const is used of same family of data type
// 2- explicit  -> by using type casting or type conversion
