package main

import "fmt"

func main() {

	var map1 map[string]string     // declaration of a map
	map1 = make(map[string]string) // instantiated

	// assign values to the map

	map1["560086"] = "Bangalore-1" // assign a value to the map
	map1["560096"] = "Bangalore-2"
	map1["560034"] = "Bangalore-3"
	map1["560000"] = "Bangalore-4"
	map1["522001"] = "Guntur-1"

	//fmt.Println(map1["560086"]) // fetch a value from the map
	for k, v := range map1 {
		fmt.Println("Key:", k, "Value:", v)
	}

}

// maps have keys and values
// value is mapped to the key
// maps are very efficient from time complexity perspective
// Big O(1)
// map is unordered. When getting values, the order of values cannot be guarrentied
// a948904f2f0f479b8f8197694b30184b --> 4
// 0d2ed1c1cd2a1ec0fb85d299a192a447 --> 3

// b948904f2f0f479b8f8197694b30188a  --> 3
// 0b8ed1c1cd2a1ec0fb85d299a192a876  --> 1
