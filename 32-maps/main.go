package main

import "fmt"

func main() {

	map1 := make(map[string]any)
	map1["India"] = 91
	map1["Pakisthan"] = 92
	map1["USA"] = 1
	map1["USA-2"] = 1
	map1["UK"] = 22

	keys := GetKeys(map1)
	fmt.Println(keys)
	values := GetValues(map1)
	fmt.Println(values)

}

func GetKeys(m map[string]any) []string {
	keys := make([]string, len(m)) // clearly know the number of keys so created a slice with the nubmer
	index := 0                     // since using range loop on a map , cannot get an index so created index
	for k := range m {             // iterating only thru keys so no need of k,v kind of
		keys[index] = k // when each key is iterated that is assigned to the slice
		index++         // increment the index so that the next key will be added to the next element of the slice
	}
	return keys // return slice with all keys in a map
}

func GetValues(m map[string]any) []any {
	keys := make([]any, len(m)) // clearly know the number of values so created a slice with the nubmer
	index := 0                  // since using range loop on a map , cannot get an index so created index
	for _, v := range m {       // iterating only thru values so no need of k,v kind of
		keys[index] = v // when each value is iterated that is assigned to the slice
		index++         // increment the index so that the next value will be added to the next element of the slice
	}
	return keys // return slice with all values in a map
}
