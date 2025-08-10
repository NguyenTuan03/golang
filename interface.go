package main

import "fmt"

func Interface() {
	fmt.Println("Hello")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("tuan nguyen"))

	myInt := IntCounter(0)
	var inc Increment = &myInt
	fmt.Print(inc.Increment())
}

type Writer interface {
	Write([]byte) (int, error)
}
type Increment interface {
	Increment() int
}

// Để dùng thì khai báo struct

type ConsoleWriter struct {
}
type IntCounter int

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
