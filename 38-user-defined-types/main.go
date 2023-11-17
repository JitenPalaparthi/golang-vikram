package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "Jitenp@outlook.com", Addr: Address{Line1: "Door no 4-43", City: "Hyderabad", State: "Telangana", Country: "India", Pincode: "522456", Pobox: 4288}, SocialMedias: []SocialMedia{{Type: "Linkedin", Handle: "linkedin.com/jpalaparthi"}, {Type: "Twitter", Handle: "Jiten_1983"}}}
	fmt.Println(p1)
	fmt.Println("Email:", p1.Email)
	fmt.Println("State:", p1.Addr.State)
	fmt.Println("Linkedin", p1.SocialMedias[0].Handle)
}

type Person struct {
	Name, Email  string
	Addr         Address       // Composition:Creating a variable of struct inside another struct
	SocialMedias []SocialMedia // slice of another user defined type
}

type Address struct {
	Line1, City, State, Country, Pincode string
	Pobox                                int
}

type SocialMedia struct {
	Type   string
	Handle string
}
