package main

// There are not public, private, protected or similar access specifiers or modifiers in Go
// The concept is very simple. 1. exported or 2. unexported
// 1. exported are similar to public
// types or fields inside types or anything those are declared in the pacakge ; if they start with Upper Case character , they are exported
// 2. unexported are similar to private
//

import (
	"fmt"
	"myshapes/shapes/rect"
	"myshapes/shapes/square"
)

func main() {
	r1 := new(rect.Rect)
	r1.L = 23.34
	r1.B = 23.545
	a1 := r1.Area()
	p1 := r1.Perimeter()
	fmt.Println("Area of rect r1:", a1)
	fmt.Println("Perimeter of rect r1:", p1)
	//r1.whoAmI() // this is unexported as it starts with lowercases
	r1.WhoAmI() // this is exported as it starts with UpperCase

	var s1 square.Square = 12.45

	a2 := s1.Area()
	p2 := s1.Perimeter()

	fmt.Println("Area of Square s1:", a2)
	fmt.Println("Perimeter of Square s1:", p2)

}
