package main

import "fmt"

func main() {

	var p1 Person // declaring a variable called p1 for Person
	// assign values to p1
	p1.Name = "JP"
	p1.Email = "Jitenp@outlook.com"
	p1.Mobile = "9618558500"
	p1.Age = 40
	p1.Address = "Door No 1-43, Guntur, AP, India"
	fmt.Println("Person p1 dettails:", p1)

	var num1 integer = 100
	fmt.Println(num1)

}

type integer = int // alias.. Where ever instead of int you can use integer

type Person struct { // creating a new type
	Name    string
	Email   string
	Mobile  string
	Age     uint8
	Address string
}

// Defined Types: Types that are created in Golang using builtin package
// User Defined Types:
