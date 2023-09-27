package main

import (
	"fmt"
	"reflect"
)

// short hand declaration of variable creation

func main() {

	var num1 = 100   // not giving a data type.
	var age = 38     // what is the data type of age : int but not int8
	var float1 = 1.4 // it is float64 not float32

	// The data type is automatically taken based on the value assigned
	// when the data type is not mentioned it takes the biggest data type from the same types

	fmt.Println("Value of num1:", num1, "Type of num1:", reflect.TypeOf(num1))
	//fmt.Println("Value of age:", age, "Type of age:", reflect.TypeOf(age))
	fmt.Printf("Value of age:%v, Type of age:%T\n", age, age)
	// fmt.Println()
	// fmt.Println()
	fmt.Println("Value of float1:", float1, "Type of float1:", reflect.TypeOf(float1))

	// short hand declartion of variables in go

	num2 := 100 // shorthand declartion

	str1 := "Hello World"

	ok1 := false

	float2 := 3.14

	fmt.Println(num2, str1, ok1, float2)

	const pi = 3.14 // this is a short handle declaration of constant. pi is float64
	fmt.Println("Value of pi:", pi, "Type of pi:", reflect.TypeOf(pi))
}
