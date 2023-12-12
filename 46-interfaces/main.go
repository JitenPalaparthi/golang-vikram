package main

import "fmt"

// Interface is a contact
// Interface contains specifications
//

type IArea interface {
	Area() float32
}

type Rect struct {
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

type Square struct {
	S float32
}

func (s *Square) Area() float32 {
	return s.S * s.S
}

func main() {
	var iArea1 IArea
	r1 := new(Rect)
	r1.L = 123.13
	r1.B = 124.56

	iArea1 = r1
	a1 := iArea1.Area() // can invoke Area method of Rect through interface
	fmt.Println("Area of Rect:", a1)

	s1 := new(Square)
	s1.S = 12.45
	iArea1 = s1 // Now iArea holds the instance of Square
	a2 := iArea1.Area()
	fmt.Println("Area of Square:", a2)
}
