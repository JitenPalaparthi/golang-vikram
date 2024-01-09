package main

import (
	"errors"
	"fmt"
)

func main() {

	str := new(string)

	*str = "Hello World"

	rstr, err := reverseStr(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*rstr)
	}

	rstr1, err := reverseStr(nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*rstr1)
	}

	str2 := new(string)
	*str2 = ""

	rstr2, err := reverseStr(str2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*rstr2)
	}
}

// funcs can return more than one value
// errors are just values
// error is an interface
func reverseStr(str *string) (*string, error) {

	// if str == nil  {
	// 	return nil, errors.New("nil input")
	// }
	// if *str == "" {
	// 	return nil, errors.New("empty input")
	// }

	if str == nil || *str == "" {
		return nil, errors.New("input is nil or empty")
	}

	rstr := new(string) // creating a string pointer

	// if str == nil {

	// }

	for _, v := range *str {
		*rstr = string(v) + *rstr
	}

	return rstr, nil
}

// error handling in Go
// errors are just values in Go
//

// try

// catch

// finally
