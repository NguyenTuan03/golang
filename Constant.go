package main

import "fmt"

const a = 100 //biến global vẫn dùng được
const (
	red = iota
	yellow
	black
)

const (
	_     = iota
	one   = 1
	two   = 2
	three = 3
)

func Constant() {
	const a = 200
	const i int = 42
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", red, red)
	fmt.Printf("%v, %T\n", yellow, yellow)
	fmt.Printf("%v, %T\n", black, black)
	fmt.Printf("%v, %T\n", one, one)
	fmt.Printf("%v, %T\n", two, two)
	fmt.Printf("%v, %T\n", three, three)
}

//tự động khởi tạo giá trị cho Enum bằng iota 0->n : type int32
