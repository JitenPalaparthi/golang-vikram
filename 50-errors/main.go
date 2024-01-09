package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("abcd.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		buf := make([]byte, 10)
		// read:
		// 	n, err := f.Read(buf)
		// 	if n > 0 || err == nil {
		// 		fmt.Print(string(buf))
		// 		goto read
		// 	}

		for {
			n, err := f.Read(buf)
			if n == 0 || err != nil {
				break
			}
			fmt.Print(string(buf))
		}
		println()
	}
}
