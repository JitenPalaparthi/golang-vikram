package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var state string = "AP"
	var gender string = "F"
	var height float32 = 3.
	var age uint8 = 5

	var input3, input4 string

	//if (state=="AP" || state=="ap") //Ap aP
	//println(strings.ToUpper(state) == "AP" && (gender == 'F' || gender == 'f') && height <= 3. && age <= 5)
	println("enter state (AP|TN|UP|DEL)")

	_, err := fmt.Scanf("%s", &state)
	if err != nil {
		println("there seems to be error in the user input so exiting-->", err.Error())
		return
	}

	println("enter gender (F|M)")

	_, err = fmt.Scanf("%s", &gender)
	if err != nil {
		println("there seems to be error in the user input so exiting-->", err.Error())
		return
	}

	println("enter height")

	_, err = fmt.Scanf("%s", &input3)
	if err != nil {
		println("there seems to be error in the user input so exiting-->", err.Error())
		return
	}
	h, err := strconv.ParseFloat(input3, 32)

	if err != nil {
		println("there seems to be error in the user input so exiting-->", err.Error())
		return
	}

	height = float32(h)
	println("enter age")

	_, err = fmt.Scanf("%s", &input4)

	if err != nil {
		println("there seems to be error in the user input so exiting-->", err.Error())
		return
	}

	a, err := strconv.Atoi(input4)

	if err != nil {
		println("there seems to be error in the user input so exiting-->", err.Error())
		return
	}

	age = uint8(a)

	println(state, gender, input3, input4)
	if strings.ToUpper(state) == "AP" && (gender == "F" || gender == "f") && height <= 3. && age <= 5 {
		println("no ticket")
	} else if gender == "F" || gender == "f" {
		println("no ticket")
	} else {
		println("there is a ticket")
	}

}

// state gender height  age ticket
// AP     M/F    <3f    <=5	no ticket
// TN     F			    	no ticket
// TN     M      <2.5f  <=6	no ticket
// Del    F             	no ticket
// Del    M      <3.5f  <=5	no ticket
// UP     M      <3f    <=6	no ticket
// UP     F      <3.5f  <=6	no ticket

// stdin
// stdout
