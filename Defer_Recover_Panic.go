package main

import "fmt"

func DeferRecoverPanic() {
	fmt.Println("Bắt đầu")

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	fmt.Println("Kết thúc")

	//Panic
	a := 10
	b := 2
	fmt.Println(a / b)
	// panic("Chia cho 1 so 0")
	// example recover

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Đã recover từ panic:", r)
		}
	}()

	fmt.Println("Chạy bình thường")
	panic("Có lỗi xảy ra")
	fmt.Println("Không in ra dòng này")

}

// defer được dùng để trì hoãn việc thực thi một hàm cho đến khi hàm bao quanh nó kết thúc (thường là trước khi return
//Khi gặp defer, Go lưu lại lời gọi hàm đó và thực thi nó sau cùng, khi hàm chứa defer kết thúc.
// Nếu có nhiều defer, chúng sẽ thực thi theo thứ tự LIFO (Last In First Out).
//Recover: một "phao cứu sinh" để bắt và xử lý panic, giúp chương trình không bị sập hoàn toàn.
//recover() chỉ hoạt động bên trong một defer.
//Nếu gọi trực tiếp bên ngoài defer, recover() luôn trả về nil.