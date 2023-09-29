package main

func main() {
	// if , else , else if
	num := 6

	println(true && true)   // true only when both or them are true then it is true
	println(true && false)  // false
	println(false && true)  // false
	println(false && false) // false

	println(true || true)   // true only when both or them are false then it is false
	println(true || false)  // true
	println(false || true)  // true
	println(false || false) // false

	println(num%2 == 0 && num%3 == 0) // true && true --> true
	// / %
	if num%2 == 0 { // if true
		println(num, "is even number")
	} else { // if false
		println(num, "is odd number")
	}

	if true {
		println("it is true")
	} else {
		println("it is false")
	}

	println("what is the ultimate result:", (true || false) && (true && true) && (2%2 == 1))
	// true && true && false
	// true && false
	// false

	if (true || false) && (true && true) && (2%2 == 1) {
		println("am i true")
	} else {
		println("am i false")
	}

}

// what is an expression? --> it yields a value
// what is a statement?  --> it executes
