// any is alias for interface{}
// rune is alias for int32
// byte is alias for uint8

package main

import (
	"fmt"
	"reflect"
)

// const (
// 	pi = 3.14
// )

// var (
// 	name = "Golang Programming"
// )

func main() {
	var any1 any
	//var any2 interface{}

	any1 = 100

	//var num1 uint8 = uint(any1) // this does not work; becase it is type assertion but not type caste

	fmt.Println("type of any1:", reflect.TypeOf(any1))
	var num1 uint8 = uint8(any1.(int))
	println("num1:", num1)

	any1 = false
	fmt.Println("type of any1:", reflect.TypeOf(any1))

	var ok1 bool
	ok1 = any1.(bool)
	fmt.Println("ok1:", ok1)

	any1 = 100.12
	any1 = "Hello World"
	any1 = 'A'
	any1 = 213123123123123

	// str1 := "Hello World by string" // in general strings are stored in heap memory but in go it depends
	// println(str1)
	// name = "Hello Golang training" + " By Jiten Palaparthi"
	// println(str1)
	fmt.Println(any1)
}

// int -> 4 or 8
// byte -> 1
// float32 -> 4

// escape analysis: it analysis that a variable is allocated in stack or in heap
// 1- moved to heap
// 2- escapes to heap

// when ever any datatype or interface{} is involved , the conversions are called as assertions

// what is called type assertion and when to use type assertion?
// difference between type assertion and type casting/ conversion
