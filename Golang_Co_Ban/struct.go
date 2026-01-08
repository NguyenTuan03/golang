package main

import (
	"fmt"
)

type hcm struct {
	Chieudai  int `json:"chieudai"`
	Chieurong int `json:"chieurong"`
}

func (h *hcm) chuvi() float32 {
	return 2 * float32(h.Chieudai+h.Chieurong)
}
func (h *hcm) dientich() int {
	return h.Chieudai * h.Chieurong
}
func Struct() {
	fmt.Println("Please enter the function you want to do")
	fmt.Println("[1] TINH CHU VI HCN")
	fmt.Println("[2] TINH DIEN TICH HCN")
	fmt.Println("[3] Exit")
	fmt.Println("============================")
	var choice int32
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("NHAP CHIEU DAI: ")
		var a int
		fmt.Scan(&a)
		fmt.Println("NHAP CHIEU RONG: ")
		var b int
		fmt.Scan(&b)
		h := hcm{
			Chieudai:  a,
			Chieurong: b,
		}
		fmt.Println("CHU VI HCN LA: ", h.chuvi())
	case 2:
		fmt.Println("NHAP CHIEU DAI: ")
		var a int
		fmt.Scan(&a)
		fmt.Println("NHAP CHIEU RONG: ")
		var b int
		fmt.Scan(&b)
		h := hcm{
			Chieudai:  a,
			Chieurong: b,
		}
		fmt.Println("DIEN TICH HCN LA: ", h.dientich())

	case 3:
		return
	default:
		fmt.Println("Please enter right function in the system")
	}
}
