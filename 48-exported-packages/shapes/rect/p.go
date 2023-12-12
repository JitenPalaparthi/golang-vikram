package rect

func (r *Rect) Perimeter() float32 {
	//return 2 * (r.l + r.b)
	return 2 * (r.L + r.B)

}
