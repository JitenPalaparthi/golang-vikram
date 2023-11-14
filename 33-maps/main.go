package main

import "fmt"

func main() {

	map1 := make(map[string]any)
	map1["India"] = 91
	map1["Pakisthan"] = 92
	map1["USA"] = 1
	map1["USA-2"] = 1
	map1["UK"] = 22
	fmt.Println(map1)
	delete(map1, "USA-2") // delete is a built in function specially used only in maps
	fmt.Println(map1)

	clear(map1)
	fmt.Println(map1)

}
