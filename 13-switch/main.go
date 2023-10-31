package main

func main() {
	num := 6
	switch {
	case num%8 == 0:
		println(num, " is divisible by 8")
		fallthrough
	//case num%6 == 0:
	//	println(num, " is divisible by 6")
	//	fallthrough
	case num%4 == 0:
		println(num, " is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, " is divisible by 2")
	}
}

//similar to break is avoided in c,c++ etc to do in go you need to add a keywowd call fallthrough
