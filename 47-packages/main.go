package main

import (
	"fmt"       // standard package
	"math/rand" // standard package
	"os"
	"time"              // standard package
	"usershapes/shapes" // user defined package. usershapes is name of the module.So used as virtual root

	yaml "gopkg.in/yaml.v2" // third party package
)

// 1.standard package: here fmt, time and rand are standard packages
// Where do standard packages reside?
// Where ever GOROOT there all standard packages reside: /usr/local/go/src (example)
//
// 2. user defined package: These are packages , created by the programmer
// How to create a user defined package
// each and every package must have a directory
// generally the name of the directory is the name of the package
// file name can be any, there can be any number of files inside a directoy for a one package
// to use user defined package, the root path is a virtual path that is the name of the module
// the name of the module is the root of the path. here usershapes is the name of the module so it is the virtual root
//
// 3. Third party package: packages are created by third parties
// search for a package to suffice your need
// download the pacakge and include into your project
// How to use: go get gopkg.in/yaml.v2
// or use go mod tidy

func main() {
	fmt.Println(time.Now())
	fmt.Println(rand.Int())

	r1 := new(shapes.Rect)
	r1.L = 123.43
	r1.B = 123.46
	a1 := r1.Area()
	fmt.Println("Area of Rect:", a1)

	data, err := os.ReadFile("data.yaml") // read yaml file into byte array
	if err != nil {
		fmt.Println(err)
	}
	mymap := make(map[string]any)     // instantiate a map
	err = yaml.Unmarshal(data, mymap) // converting read byte array into map
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mymap) // just print the maps
}
