package main

func main() {
	num := 55
	switch { // empty switch
	case num >= 0 && num <= 50:
		println(num, "between or equal to 0-50")
	case num > 50 && num <= 100:
		println(num, "between greater than 50 and less than = 100")
	default:
		println(num, " seems to be greater than 100 or less than 0")
	}
}

// c, c++, c#,java
