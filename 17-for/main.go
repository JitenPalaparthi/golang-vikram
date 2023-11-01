package main

func main() {

outer: // lable
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == 6 {
				break outer
			}
			println("i:", i, "-->", "j:", j)
			if j == 5 {
				break
			}
		}
	}
}

// identifiers --> should not start with number, should not have a specialchar other than _. cannot have space between variable
// Big O
// time vs space
// n * n
