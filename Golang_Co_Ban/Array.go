package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println("Gia tri arr[", i, "] =", arr[i])
	}
	arrString := [3]string{"Golang", "Java", "Python"}
	for i := 0; i < len(arrString); i++ {
		fmt.Println("Gia tri arrString[", i, "] =", arrString[i])
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
}
