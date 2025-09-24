package main

import "fmt"

func main() {
	//B1
	// for i := 1; i <= 100; i++ {
	// 	if (i == 6 || i == 48 || i == 75 || i == 89) {
	// 		i++;
	// 	}
	// 	if (i == 100) {
	// 		fmt.Print(i, ".")
	// 		break;
	// 	}
	// 	fmt.Print(i, ", ")
	// }
	//B2
	for i := 1; i <= 100; i += 2 {
		for j := 1; j <= 3; j++ {
			fmt.Print(i, ", ")
		}
		fmt.Println()
	}
}
