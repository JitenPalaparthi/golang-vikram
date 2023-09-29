package main

import (
	"fmt"
	"math"
)

func main() {

	var num1 uint8 = 100

	var num2 uint16 = 1231

	var num3 uint32 = 123123

	var num4 uint64 = 11213123131

	var num5 int8 = -100

	var num6 int16 = -12321

	var num7 int32 = -231313

	var num8 int64 = -12312312312312

	var num9 int = -3213131

	var num10 uint = 31231231313123

	var float1 float32 = 123213.123

	var float2 float64 = 1212332.12312331

	var any1 any = 13433.343 // float64

	var any2 any = 1323123 // int

	var sum float64

	sum = float64(num1) + float64(num2) + float64(num3) + float64(num4) + float64(num5) + float64(num6) + float64(num7) + float64(num8) + float64(num9) + float64(num10) + float64(float1) + float2 + any1.(float64) + float64(any2.(int))

	fmt.Println("Sum:", sum)
	fmt.Printf("Sum:%.2f\n", sum)
	fmt.Printf("Sum ceil:%f\n", math.Ceil(sum))
	fmt.Printf("Sum floor:%f\n", math.Floor(sum))
	fmt.Printf("Sum abs:%f\n", math.Abs(sum))
	//fmt.Println("Sum ceil:", math.Ceil(sum))
	// format verbs in go for Printf
}
