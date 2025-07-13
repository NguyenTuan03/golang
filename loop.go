package main

import "fmt"

func Loop() {
	for i := 1; i <= 10; i++ {
		if (i%2 ==0) {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println();
	//exmaple 2
	a := []int{5,8,3,9}
	for i,value := range a {
		fmt.Println(i, value)
	}
	fmt.Println();
	//example 3
	str := "Hello world"
	for index, str := range str {
		fmt.Println(index, string(str))
	}
}