package main

import "fmt"

type Employee struct {
	ID    int    `json:Id`
	Name  string `json:Name`
	phone int `json:Phone`
}

func Array() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println("Gia tri arr[", i, "] =", arr[i])
	}
	arrString := [3]string{"Golang", "Java", "Python"}
	for i := 0; i < len(arrString); i++ {
		fmt.Println("Gia tri arrString[", i, "] =", arrString[i])
	}
	for i, val := range arr {
		fmt.Println(i, val)
	}
	// matrix
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}
	// ======================================================================
	employee := [...]Employee{
		{
			ID: 1,
			Name: "Tuan",
			phone: 1234,
		},
		{
			ID: 2,
			Name: "TBinuan",
			phone: 134,
		},
		{
			ID: 1,
			Name: "OKOK",
			phone: 1,
		},
	}
	for _, val := range employee {
		fmt.Printf("ID: %d, Name: %s, phone: %d\n", val.ID, val.Name, val.phone)
	}
}
