package main

import "fmt"

func main() {

	var c1 ColourCode
	c1.int = 100
	c1.string = "Red"
	c1.string1 = "Is it really Red?"
	fmt.Println(c1)

}

// Struct with anonymous fields
// one of the disadvantages is cannot have multiple ints or multiple strings.
// filed type must be distinct
type string1 = string
type ColourCode struct {
	int
	string
	string1
}
