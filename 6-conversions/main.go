package main

import (
	"fmt"
	"strconv"
)

type integer = int // type alias

func main() {

	var num1 integer = 100
	println(num1)

	str1 := "Hello World!" // ascii chars

	str1 = "Hello Folks-->!" // since string is immutable , how can this statement is valid?

	str2 := "你好世界" // unicode chars

	fmt.Println(str1, str2)

	var r1 rune = 'H' // a char is nothing but its unicode value 72

	var b1 byte = 'H' // alias for uint8 range 0-255

	var r2 int32 = '你'

	//var b2 byte = '你' // byte is only one byte but this chinees char is an unicode char which is 20320 that is bigger than the range of byte/uint8

	// clang each char is 1 byte
	// according to golan 72 == 'H'
	// each char is 4 bytes ->

	fmt.Println(r1, b1, r2)
	fmt.Println(string(r1), string(b1), string(r2))

	// conversions in strings or chars

	var str3 = "H" //

	var num2 uint8 = str3[0] // will learn more abt index in arrays

	//var num3 uint8 = str3

	var str4 = "2023" // there are 4 chars 2 0 2 3

	var num3 int = int(str4[0])

	var num4 int
	num4, _ = strconv.Atoi(str4) // as long as the data in a string is int then it converts, otehrwise it returns err
	// _ blank identifier : a function returns a value but you dont want to use it then use blank identifer

	num5 := 2023
	var str5 string = strconv.Itoa(num5)
	var str6 string = fmt.Sprint(num5)

	var ok = true // ok is boolen. The value of ok is either true or false
	// how to convert it to string

	var str7 string = fmt.Sprint(ok)

	fmt.Println(num2, num3, num4, str5, str6, str7)

}

// string
// an array of bytes/ chars
// string is immutable in go
// string is immutable from the internal design and memory allocation perspective but not from the
// users or developers perspective
