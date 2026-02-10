package main

import (
	"fmt"
)

type Employeee struct {
	Name string
	Age  int
	Role string
}

func Map() {
	drink := map[string]int{
		"drink":  300,
		"coffee": 500,
	}
	fmt.Println(drink["drink"])
	// Cách khai báo và khởi tạo map
	// Sử dụng make
	m := make(map[string]int)
	m["apple"] = 100
	m["organge"] = 2200
	// Khai báo và khởi tạo trực tiếp
	person := map[string]string{
		"Name":    "tuan",
		"age":     "12",
		"address": "HCM",
	}
	fmt.Println(person)
	// Khai báo rỗng
	var emptyValue map[string]int
	// Cần cấp phát bộ nhớ trước khi sử dụng
	emptyValue = make(map[string]int)
	emptyValue["apple"] = 10
	fmt.Println(emptyValue)
	//=======================================
	newperson, ok := person["ok"]
	if ok {
		fmt.Println(newperson)
	} else {
		fmt.Println("Ko ton tai")
	}
	// Map and struct
	employees := map[string]Employeee{
		"e1" : {
			Name: "Tuan",
			Age: 32,
			Role: "IT",
		},
		"e2" : {
			Name: "Binh",
			Age: 18,
			Role: "ITDASD",
		},	
	}
	for key, value := range employees {
		fmt.Println("Nguoi thu ", key)
		fmt.Println(value.Name)
		fmt.Println(value.Age)
		fmt.Println(value.Role)
	}
	// Map and slice
	studentsObject := map[string][]string{
		"Tuan": {
			"Math",
			"English",
		},
		"Thanh": {
			"1",
			"2",
		},
	}
	studentsObject["Tuan"] = append(studentsObject["Tuan"], "Co")
	for key, val := range studentsObject {
		fmt.Println("Mon cua ban ", key, " la: ")
		fmt.Println(val)
	}
}
