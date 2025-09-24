package main

import "fmt"

func main() {
	maxx := 100
	maxxOdd := 1
	count := 0
	if maxx%2 == 1 {
		maxxOdd = maxx
	} else {
		maxxOdd = maxx - 1
	}
	fmt.Println("maxxOdd:", maxxOdd)
	for i := 1; i <= maxx; i++ {
		if i%2 == 1 {
			count++
			if i == maxxOdd {
				fmt.Print(i, ".")
				break
			}
			if count == 3 {
				fmt.Print(i, "")
			} else {
				fmt.Print(i, ", ")
			}
		}
		if count == 3 {
			count = 0
			fmt.Println()
		}
	}
}
