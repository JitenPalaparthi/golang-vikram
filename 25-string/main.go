package main

import "fmt"

//type integer = int

func main() {
	str := "Hello World"
	// each char is a unicode char
	// 0-255 : ASCII char set. I byte
	// 0-255

	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	str = "హలో உலகம் folks!" // unicode charz
	// for loop iterates only byte by byte but each char in string is not one byte
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	// range loop on a string is capable of understaning each char
	for _, v := range str {
		fmt.Println(string(v))
	}
}

// string  : Is string mutable?
