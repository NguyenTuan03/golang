package main

import "fmt"

func BT() {
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
	// maxx := 100
	// maxxOdd := 1
	// count := 0
	// if maxx%2 == 1 {
	// 	maxxOdd = maxx
	// } else {
	// 	maxxOdd = maxx - 1
	// }
	// fmt.Println("maxxOdd:", maxxOdd)
	// for i := 1; i <= maxx; i++ {
	// 	if i%2 == 1 {
	// 		count++
	// 		if i == maxxOdd {
	// 			fmt.Print(i, ".")
	// 			break
	// 		}
	// 		if count == 3 {
	// 			fmt.Print(i, "")
	// 		} else {
	// 			fmt.Print(i, ", ")
	// 		}
	// 	}
	// 	if count == 3 {
	// 		count = 0
	// 		fmt.Println()
	// 	}
	// }
	//Ex3
	var start int
	var end int
	fmt.Print("Enter start: ")
	_, err := fmt.Scanf("%d", &start)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Print("Enter end: ")
	_, err = fmt.Scanf("%d", &end)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if (start < 0 || end < 0) || (start > end) || end > 10 {
		fmt.Println("Invalid input")
		return
	}
	for i := start; i <= end; i++ {
		if i%2 == 1 {
			fmt.Print(i, ", ")
		}
	}

}
