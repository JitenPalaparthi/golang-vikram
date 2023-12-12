package rect

import "fmt"

// type rect struct{ // rect is unexported since it starts with lower case
// l, b float32 // these are unexported or restricted, as l and b start with lower case
// }
type Rect struct {
	// l, b float32 // these are unexported or restricted, as l and b start with lower case
	L, B float32 // these are exported , as L and B start with Upper case
}

func (r *Rect) whoAmI() {
	//fmt.Println("I am Rect object.Length:", r.l, "Bredth:", r.b)

	fmt.Println("I am Rect object.Length:", r.L, "Bredth:", r.B)
}

func (r *Rect) WhoAmI() {
	r.whoAmI() // even it starts with lowercase it can be called internally since it is part of the pacakge and calling in the same package
}
