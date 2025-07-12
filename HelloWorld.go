package main

import (
	"fmt"
	"strconv"
)

var (
	ExportOutSide int     = 10000 // khai báo biến ngoài hàm main, chữ cái đầu tiên phải viết hoa
	n             int     = 10000 // khai báo biến ngoài hàm main
	m             string  = "My name is John"
	age           int     = 20
	pi            float64 = 3.14
	isStudent     bool    = true
)

func HelloWorld() {
	var name string = "John"
	//Có 3 cách khai báo biến Interger
	var i int
	i = 10
	name = "Jane"
	fmt.Println("Hello World", name)
	fmt.Println(i)
	n := "Nguyen"
	fmt.Println(n)
	fmt.Printf("%v, %T\n", name, n)
	var convertTypeToFloat float64 = float64(i)
	fmt.Printf("%v, %T\n", convertTypeToFloat, convertTypeToFloat)
	var convertTypeToString string = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", convertTypeToString, convertTypeToString)

}

//Shadow: bóng -> biến trong scope sẽ override biến ngoài scope
//export global scope ->
//Naming convention -> camelCase
//Convert type ->
