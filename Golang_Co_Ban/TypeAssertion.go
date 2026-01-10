package main
import "fmt"
// TypeAssertion: Kiểm tra và chuyển đổi kiểu dữ liệu của interface{} bằng cách thêm dấu .(Type)
func TypeAssertion() {
	var i any = "hello"
	if s, ok := i.(string); ok {
		fmt.Println("Gia tri laf string", s, ok)
	}

	f, ok := i.(float64)
	if !ok {
		fmt.Println("Khong phai float64", f, ok)
	}

}