package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(num1)
	str1 := fmt.Sprint(num1)
	fmt.Println(str1, reflect.TypeOf(str1))

	str2 := num2.ToString()
	fmt.Println(str2, reflect.TypeOf(str2))

	fmt.Println(num2.Square())

	sq1 := integer(num1).Square()
	fmt.Println("Square num1:", sq1)
	str3 := integer(num1).ToString()
	fmt.Println(str3, reflect.TypeOf(str3))

}

var num1 int = 100 // global variable
var num2 integer = 200

type integer int // one of the compelling features of Go. Can create a user dfined type from builtin type

// for a user defined type , can create methods

func (i integer) ToString() string {
	return fmt.Sprint(i)
}

func (i integer) Square() int {
	return int(i * i)
}
