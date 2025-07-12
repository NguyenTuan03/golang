package main

import (
	"fmt"
)

func Array() {
	arrayStPoint := []string{"10", "2", "3"} //slices: a new slice variable points to the same underlying array.
	// fmt.Printf("%v\n", len(arrayStPoint)) // length of array
	addArray := arrayStPoint
	addArray[1] = "okok"
	fmt.Println(arrayStPoint)
	fmt.Println(cap(arrayStPoint))
	fmt.Println(addArray)
	fmt.Println("Slice Formulas.......")
	//slice
	a := [...]int{2, 5, 10, 6, 7, 32, 9}
	b := a[:]  //nghĩa là muốn thay đổi con trỏ đầu và cuối -> lấy hết tất cả
	c := a[3:] // bắt đầu từ vị trí index: 3 -> cuối
	d := a[:6] // đầu -> index 6
	e := a[3:6]
	fmt.Printf("a = %v, %v, %v\n", a, len(a), cap(a))
	fmt.Printf("b = %v, %v, %v\n", b, len(b), cap(b))
	fmt.Printf("c = %v, %v, %v\n", c, len(c), cap(c))
	fmt.Printf("d = %v, %v, %v\n", d, len(d), cap(d))
	fmt.Printf("e = %v, %v, %v\n", e, len(e), cap(e))
	//make: khai báo 1 slice và có thể xác định được len và cap
	fmt.Println("Make Formulas.......")
	amake := make([]int, 0)
	amake = append(amake, 10) // dùng cho slice -> vào cuối ptu
	amake = append(amake, 2) // dùng cho slice -> vào cuối ptu
	amake = append(amake, 5) // dùng cho slice -> vào cuối ptu cap = 4
	amake = append(amake, 200) // dùng cho slice -> vào cuối ptu
	amake = append(amake, []int{4,7,3,9}...)
	fmt.Printf("%v , %T, %v, %v\n", amake, amake, len(amake), cap(amake))
	//delete the middle value
	total := []int{1,2,3,4,5,6}
	deleteMiddle := append(total[:2], total[3:]...)
	deleteMiddle = deleteMiddle[:len(deleteMiddle) - 1]
	fmt.Printf("%v, %T", total, total)
	fmt.Printf("%v, %T", deleteMiddle, deleteMiddle)
}