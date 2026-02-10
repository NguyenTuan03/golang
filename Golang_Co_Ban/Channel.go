package main

import "fmt"

func main() {
	// Channel là 1 kênh để truyền dữ liệu giữa các goroutine
	// Các loại Channel
	// Unbuffered Channel: Chỉ chứa được 1 gtri tại 1 thời điểm
	// Gửi (ch <- 1) sẽ bị chặn cho đến khi có goroutine nhận
	// Nhận (<-ch) sẽ bị chặn cho đến khi có goroutine gửi

	// Buffered Channel: Có thể chứa nhiều gtri
	// Gửi (ch <- 1) sẽ bị chặn khi buffer đầy
	// Nhận (<-ch) sẽ bị chặn khi buffer rỗng

	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

}
