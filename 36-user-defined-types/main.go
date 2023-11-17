package main

import "fmt"

func main() {

	var p0 Person // declaring a variable called p1 for Person
	// assign values to p1
	p0.Name = "JP"
	p0.Email = "Jitenp@outlook.com"
	p0.Mobile = "9618558500"
	p0.Age = 40
	p0.Address = "Door No 1-43, Guntur, AP, India"
	fmt.Println(p0)

	var p1 Person = Person{Name: "Vikram", Email: "Vikram@gmail.com", Mobile: "1234343", Age: 40, Address: "Doro no 1-1,USA"}
	fmt.Println(p1)

	// shorthand declaration
	p2 := Person{Name: "Vikram", Email: "Vikram@gmail.com", Mobile: "1234343", Age: 40, Address: "Doro no 1-1,USA"}
	fmt.Println(p2)

	// another way of shorthand declaration
	var p3 = Person{Name: "Vikram", Email: "Vikram@gmail.com", Mobile: "1234343", Age: 40, Address: "Doro no 1-1,USA"}
	fmt.Println(p3)

	// just a declaration, if no assigment then type inference works
	var p4 Person
	fmt.Println(p4) // type inference work here {   0 }
}

type Person struct { // creating a new type
	Name    string
	Email   string
	Mobile  string
	Age     uint8
	Address string
}
