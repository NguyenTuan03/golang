package main

import "fmt"

func Primitive_Data() {
	var a int8 = 10
	var b int8 = 3
	fmt.Printf("%v, %T\n", a+b, b+b)
	fmt.Printf("%v, %T\n", a-b, a-b)   // and
	fmt.Printf("%v, %T\n", a&b, a&b)   // or
	fmt.Printf("%v, %T\n", a^b, a^b)   // nor
	fmt.Printf("%v, %T\n", a&^b, a&^b) // xor
	var floata float32 = 10.4
	var floatb float64 = 2.5
	fmt.Printf("%v , %T\n", floata, floata)
	fmt.Printf("%v , %T\n", floatb, floatb)
	text := "hello world"
	fmt.Printf("%v , %T\n", text, text)
	fmt.Println([]byte(text))
	
}

//boolean
//Numeric
//Text
