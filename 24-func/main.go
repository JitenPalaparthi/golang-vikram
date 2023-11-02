package main

import "fmt"

func main() {

	sum1 := SumOf([10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19})
	fmt.Println("Sum of arr:", sum1)

	arr1 := [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	sum2 := SumOf(arr1)

	fmt.Println("Sum of arr:", sum2)

	arr2 := [3]int{213, 43, 34}
	sum3 := SumOf3(arr2)
	fmt.Println("Sum of arr:", sum3)
}

func SumOf(arr [10]int) int {
	sum := 0
	for _, v := range arr {
		//sum=sum+v
		sum += v
	}
	return sum
}

func SumOf3(arr [3]int) int {
	sum := 0
	for _, v := range arr {
		//sum=sum+v
		sum += v
	}
	return sum
}

func SumOf4(arr [4]int) int {
	sum := 0
	for _, v := range arr {
		//sum=sum+v
		sum += v
	}
	return sum
}
