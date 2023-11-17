package main

import "fmt"

func main() {

	p1 := Person{Name: "jiten", Email: "JitenP@outlook.com", Address: struct {
		Pobox int
		Line  string
	}{Pobox: 2323, Line: "Guntur,  AP,India"}}
	fmt.Println(p1)
	fmt.Println("Pobox:", p1.Address.Pobox)

	// anonymous struct
	e1 := struct { // e1 is a variable.. an anonymous struct was created and also assigned values
		Name, Email string
		Id          int
	}{Name: "Jiten", Email: "JitenP@Outlook.com", Id: 1001}

	fmt.Println(e1)
}

type Person struct {
	Name, Email string
	Address     struct { // embedding a struct in another struct Or nested struct
		Pobox int
		Line  string
	}
}
