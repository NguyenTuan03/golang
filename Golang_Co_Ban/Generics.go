package main

import (
	"cmp"
	"fmt"
	"golang.org/x/exp/constraints"
)

// Generics: là 1 tính năng cho phép định nghĩa các hàm, kiểu dữ liệu có thể hoạt động với nhiều kiểu dữ liệu khác nhau mà không cần phải viết lại mã cho từng kiểu cụ thể.
// Cú pháp sử dụng dấu [] để định nghĩa kiểu tổng quát.
// Type parameter: là một biến đại diện cho một kiểu dữ liệu cụ thể trong định nghĩa của hàm hoặc kiểu tổng quát. Nó được đặt trong dấu [] sau tên hàm hoặc kiểu.
// Type constraint: là một ràng buộc được sử dụng để giới hạn các kiểu dữ liệu mà một kiểu tổng quát có thể chấp nhận. Ràng buộc này được định nghĩa bằng cách sử dụng từ khóa "any" hoặc "comparable" hoặc bằng cách định nghĩa các interface tùy chỉnh.
// any: đại diện cho bất kỳ kiểu dữ liệu nào.
// comparable: đại diện cho các kiểu dữ liệu có thể so sánh được (các kiểu dữ liệu hỗ trợ các phép so sánh như == và !=).
func PrintValue[Type any](value Type) Type {
	return value
}
func isEqual[T comparable](a, b T) bool {
	return a == b
}
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
type Box[T any] struct {
	Content T
}
type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex // Sử dụng toán tử | để kết hợp nhiều ràng buộc thành một ràng buộc duy nhất.
}
func Sum[T Number](a, b T) T {
	return a + b
}
func Generics() {
	fmt.Println(PrintValue("Tuan"))
	fmt.Println(PrintValue(123))
	fmt.Println(PrintValue(true))
	fmt.Println("================")
	fmt.Println(isEqual(10, 20))
	fmt.Println(Max(10, 20))
	fmt.Println("================")
	box1 := Box[string]{Content: "Hello Generics"}
	box2 := Box[int]{Content: 100}
	fmt.Println(box1.Content)
	fmt.Println(box2.Content)
	fmt.Println("================")
	fmt.Println(Sum(10, 20))
}
