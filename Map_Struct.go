package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name    string
	age     int
	address string
}
type Person struct {
	name string
	age int
}
type Employee struct {
	Person,
	job string
	salary int `Maximum can't over 100VND`
} // sẽ return về string này sau khi validate
func MapStruct() {
	studentName := make(map[string]int)
	studentName = map[string]int{
		"Tuan": 20,
		"John": 24,
		"huy":  12,
	}
	studentName["Toan"] = 100 // Thêm 1 value mới
	delete(studentName, "Toan")
	_, isExist := studentName["Toan"]
	copyMap := studentName
	fmt.Println(studentName)
	fmt.Println(isExist)
	fmt.Println(copyMap)
	// C1
	// Person1 := Student {
	// 	name: "Tuan",
	// 	age: 18,
	// 	address: "HCM",
	// }
	// C2
	Person1 := Student{}
	Person1.name = "Tuan"
	Person1.address = "HCM"
	Person1.age = 20
	fmt.Println(Person1)
	Car := struct {name string}{name:"BMW"}
	copyCar := Car
	copyCar.name = "Merc"
	fmt.Println(Car)
	fmt.Println(copyCar)
	//tag trong struct
	Employee1 := Employee{}
	t := reflect.TypeOf(Employee1)
	field, _ := t.FieldByName("salary")
	fmt.Println(field)
}
