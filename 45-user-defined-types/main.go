package main

import "fmt"

func main() {
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["bangalore-1"] = "560086"
	map1["bangalore-2"] = "560096"
	map1["bangalore-3"] = "560034"

	var map2 MapString

	map2 = make(MapString)

	map2["bangalore-1"] = "560086"
	map2["bangalore-2"] = "560096"
	map2["bangalore-3"] = "560034"

	keys1 := map2.GetKeys()
	vals1 := map2.GetValues()
	fmt.Println("Keys:", keys1)
	fmt.Println("Vals:", vals1)

	keys2 := MapString(map1).GetKeys()
	vals2 := map2.GetValues()
	fmt.Println("Keys:", keys2)
	fmt.Println("Vals:", vals2)

}

type MapString map[string]string // User defined type from the underlining type

func (m MapString) GetKeys() []string {
	keys := make([]string, len(m))
	index := 0
	for k := range m {
		keys[index] = k
		index++
	}
	return keys
}

func (m MapString) GetValues() []string {
	vals := make([]string, len(m))
	index := 0
	for _, v := range m {
		vals[index] = v
		index++
	}
	return vals
}
