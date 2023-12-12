package main

import (
	"fmt"
)

func main() {
	greet()

	r1 := &Rect{L: 12.34, B: 12.45}
	a1 := Area(*r1)
	a2 := r1.Area() // To call a method
	p2 := r1.Perimeter()
	fmt.Println("Area of Rect r1:", a1)
	fmt.Println("Area of Rect r1:", a2)
	fmt.Println("Perimeter of Rect r1:", p2)

	fmt.Println("Area that is stored in the Rect r1:", r1.A)
	fmt.Println("Perimeter that is stored in the Rect r1:", r1.P)

}

// What is a method?
// How methods are different from functions?

func greet() {
	fmt.Println("Hello World!")
}

type Rect struct {
	L, B float32
	A, P float32
}

// Area is a function, not a method
func Area(r1 Rect) float32 {
	return r1.L * r1.B
}

// Area is a method
func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

// Perimeter is a methods
func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}
