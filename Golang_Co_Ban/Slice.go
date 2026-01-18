package main

import (
	"fmt"
	"reflect"
	"slices"
)

// slice: 1 tập hợp các giá trị có cùng kiểu dữ liệu nhưng kích thước có thể thay đổi linh hoạt
// Cách khởi tạo slice
// 1. Sử dụng cú pháp literal ví dụ:
//    slice := []int{1, 2, 3, 4, 5}
// 2. Sử dụng hàm make
// 3. Slicing từ mảng
func main() {
	var nums [3]int = [3]int{1,2,3}
	fmt.Println("check type ", reflect.TypeOf(nums).Kind())
	fmt.Println("check type ", reflect.Slice)
	// slice có 3 phần từ và dung lượng tối đa là 5
	slice := make([]int, 3, 5)
	slice = append(slice, 1, 2, 3, 4, 5, 5, 6, 7)
	// fmt.Println("slice ", slice)
	// fmt.Println("slice ", cap(slice))
	// for i,val := range slice {
	// 	fmt.Printf("index: %d, value: %d\n", i, val)
	// }
	// slice đa chiều
	school := [][]string{
		{"Toan", "Ly", "Hoa"},
		{"Su", "Dia", "Van"},
		{"Anh", "Phap", "Duc"},
	}
	for i,val := range school {
		fmt.Printf("Lop %d: %v\n", i+1, val)
		for j, mon := range val {
			fmt.Printf("  Mon %d: %s\n", j+1, mon)
		}
	}
	fmt.Println(slice)
	// Built-in functions for slices
	copied := slices.Clone(slice)
	fmt.Println("Copied slice:", copied)

	compareSlices := slices.Equal(copied, slice)
	fmt.Println("Are slices equal?", compareSlices)

	findIndexByValueSlice := slices.Index(slice, 3)
	fmt.Println("Index of value 3 in slice:", findIndexByValueSlice)

	findValueInSlice := slices.Contains(slice, 6) // return bool
	fmt.Println("Does slice contain value 6?", findValueInSlice)

	insertValueIntoSlice := slices.Insert(slice, 2, 99, 100)
	fmt.Println("Slice after insertion:", insertValueIntoSlice)

	deleteAnyValueFromSlice := slices.Delete(slice, 0, 4) // delete from index 1 to index 4-1
	fmt.Println("Slice after deletion:", deleteAnyValueFromSlice)
}