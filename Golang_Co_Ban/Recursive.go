package main

import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}
func listfibonacci(n int) {

	term1, term2 := 0, 1
	if n == 1 {
		fmt.Print(term1)
	} else if n == 2 {
		fmt.Print(term1, term2)
	} else {
		fmt.Print(term1, " ", term2, " ")
		for i := 3; i <= n; i++ {
			term3 := term1 + term2
			fmt.Print(term3, " ")
			term1 = term2
			term2 = term3
		}
	}

}
func sumFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else if n == 0 {
		return 0
	}
	return sumFibonacci(n-1) + sumFibonacci(n-2)
}
func main() {
	fmt.Println("Please enter the function you want to do")
	fmt.Println("[1] Sum from 1 to n")
	fmt.Println("[2] List fibonacci")
	fmt.Println("[3] Sum fibonacci")
	fmt.Println("[4] Exit")
	fmt.Println("============================")
	var choice int32
	fmt.Scan(&choice)
	switch choice {
	case 1:

		var numb int
		fmt.Println("Enter a number: ")
		_, err := fmt.Scan(&numb)
		if err != nil || numb < 0 {
			fmt.Println("A number must be a positive number")
		}
		fmt.Println(sum(numb))

	case 2:
		var numb int
		fmt.Println("Enter a number: ")
		_, err := fmt.Scan(&numb)
		if err != nil || numb < 0 {
			fmt.Println("A number must be a positive number")
		}
		listfibonacci(numb)
	case 3:
		var numb int
		fmt.Println("Enter a number: ")
		_, err := fmt.Scan(&numb)
		if err != nil || numb < 0 {
			fmt.Println("A number must be a positive number")
		}
		fmt.Println(sumFibonacci(numb))
	case 4:
		return
	default:
		fmt.Println("Please enter right function in the system")
	}

}
