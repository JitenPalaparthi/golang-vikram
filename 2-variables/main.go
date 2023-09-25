package main

func main() {
	println("Hello World!")
	var age uint8  // it needs 1 byte  ----------------------------------------------/
	var num1 int   // 8 bytes depends on arch is 32 bit 4 bytes or 64 bit 8 bytes
	var num2 int16 // 2 bytes -16k +16k
	var num3 int32 // 4 bytes -32k + 32k
	var num4 int64 // 8 bytes
	var num5 int8  // -128 + 128

	{ //-----------/
		var num6 int8 = -120
		println(num6)
	} //-----------/

	var float1 float32
	var float2 float64

	var ok bool // 1 byte .. type inference infers false by default

	var str1 string                 // default value is empty string ""
	var str2 string = "Hello World" // default value is empty string ""

	println(age, num1, num2, num3, num4, num5, float1, float2, ok, str1, str2)

	// allocations and deallocations
	// stack pointer
	//  --    sp
	//  --   stack frame -1
	//  --   			|-- num6
	//  --   num5
	//  --   num4
	//  --   num3
	//  --   num2
	//  --   num1
	//  --   age

} //---------------------------------------------------------------------------------/

// variables
// numbers
// int,int8,int16,int32,int64,uint8,uint16,uint32,uint64
// float32, float64

// when an appliction is run
// a process is created  by os
// some memory is given to the process by the os
// a thread is by default created in the process by the os
// the memory is divided into few parts
// code segment, data segment, stack and heap
// go run main.go
// compiles + links + assembles + assembler --> generates a binary --> execute
//
