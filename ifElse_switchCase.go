package main

import "fmt"

func main() {
	//Example 1:
	m := make(map[string]int)
	m = map[string]int{
		"Tuan": 19,
		"John": 20,
	}
	if age, isExist := m["Tuan"]; isExist { // tạo ra biến age, và optional boolean để check
		fmt.Println(age)
	} else {
		fmt.Println("Wrong")
	}
	//example 2:
	guess := 1
	switch guess {
	case 1:
		fmt.Println("Green")
	case 2:
		fmt.Println("REd")
	default: 
		fmt.Println("default blue")
	}
	//example 3:
	var i interface{} = 1;
	switch i.(type) {
		case int:
			fmt.Println("Type integer")
		case float32: 
			fmt.Println("Type float")
		default:
			fmt.Println("No Type!")
	}
}
