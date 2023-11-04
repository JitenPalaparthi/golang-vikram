package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 10)
	for i, _ := range slice {
		slice[i] = rand.Intn(100)
	}

	// slice2 := slice1 // is it a deep copy or shallow copy
	// i := 10
	// j := i // deep copy

	{
		{
			sum, sub, mul := calc(slice)
			fmt.Println(sum, sub, mul)
			{
				var s1, sb1, m1 int
				{
					s1, sb1, m1 = calc(slice)
					fmt.Println(s1, sb1, m1)
				}
				//	fmt.Println(s1) // you can access s1 here becase the scope is till next line
			}
		}
	}
}

// := creating a new variable and assign
// = using the declared variable and assign or mutate

// calc func takes a slice as a input argument and
// it returns three integers.1st one is sum of slice. 2nd one is substraction of slice
// the 3rd one is multiplication is slice.
func calc(s []int) (sum int, sub int, mul int) {
	//sum, sub, mul := 0, 0, 1 // creation and assign values
	if mul == 0 { // no point is checking mul is 0 or not(by default num is 0) but to explain scope.
		mul = 1
	}
	for _, v := range s { // v is created inside for-range scope.
		sum += v // add = add+v // mutate
		sub -= v // mutate
		mul *= v // mutate
	} // v cannot be used beyond this point
	// fmt.Println(v)
	return sum, sub, mul
}

// What is scope ? not but a boundary
// How to declare scope? : {}
// How many scopes in a program? any number
// Can I create a scope without any purpose? Yes
// where do you create scope? Where ever you see {=open }=close mustach brackets it is scope
// what happens when variables are created in a scope? in most of the cases they live as long as the scope lives. But there are some rules (will discuss later)
// what is a stack frame? a stack frame is created for the variable declarations. For a stack frame variables may be created or referenced.
// how does it work for shallow copy variables and deep copy variables?
