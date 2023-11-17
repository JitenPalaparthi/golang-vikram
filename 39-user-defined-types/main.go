package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "Jitenp@outlook.com", Status: "Active", Address: Address{Line1: "Door no 4-43", City: "Hyderabad", State: "Telangana", Country: "India", Pincode: "522456", Pobox: 4288, Status: "Living"}, SocialMedias: []SocialMedia{{Type: "Linkedin", Handle: "linkedin.com/jpalaparthi"}, {Type: "Twitter", Handle: "Jiten_1983"}}}
	fmt.Println(p1)
	fmt.Println("Email:", p1.Email)
	//fmt.Println("State:", p1.Address.State)
	fmt.Println("State:", p1.State) // can call state directly ,which is from Address becasue Address is used as a promoted field

	fmt.Println("Status of Person:", p1.Status)
	fmt.Println("Status of Person's Address:", p1.Address.Status) // Though calling the status of Address , since Status is in both Address and Person , need to mention Address.Status

	// can also assign values directly for Promoted fields

	p1.Pobox = 54999 // Pobox is from Address but Address is promoted field, can call directly with p1

	fmt.Println("Pobox", p1.Pobox)
	fmt.Println("Linkedin", p1.SocialMedias[0].Handle)
}

// using Address as a promoted field
// promoted field is using the userdefined type as a field
type Person struct {
	Name, Email, Status string
	Address                           // promoted filed
	SocialMedias        []SocialMedia // slice of another user defined type
}

type Address struct {
	Line1, City, State, Country, Pincode, Status string
	Pobox                                        int
}

type SocialMedia struct {
	Type   string
	Handle string
}
