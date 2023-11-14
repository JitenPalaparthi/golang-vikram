package main

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()

	i := 1 // shorthand declaration of a variable called i
loop: // label here it is loop but you can give any valid identifier as a label
	print(i, " ") // print i
	i++           // increment
	if i <= 10 {  // check the condition if the the condition is true then
		goto loop // go to the label called loop. If not satisfied go the next line
	}
	println("\nThe goto ends. So I reached here")
}

// goto can be used to jump to the different portions of the code.
// the label is important. Label name can be of any valid identifier
