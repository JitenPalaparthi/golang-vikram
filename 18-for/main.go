package main

func main() {

	// multiple conditions and initis

	for i, j := 1, 2; i <= 4 && j <= 3; i, j = i+1, j+1 {
		println(i, "--", j)
	}

}

// &&
