package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func task(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Task", id, "is running")
	time.Sleep(1 * time.Second)
	fmt.Println("Task", id, "is ending")
	ch <- "Task " + strconv.Itoa(id) + " is done"
}
func Channel() {
	// Channel là 1 kênh để truyền dữ liệu giữa các goroutine
	// Các loại Channel
	// Unbuffered Channel: Chỉ chứa được 1 gtri tại 1 thời điểm
	// Gửi (ch <- 1) sẽ bị chặn cho đến khi có goroutine nhận
	// Nhận (<-ch) sẽ bị chặn cho đến khi có goroutine gửi

	// Buffered Channel: Có thể chứa nhiều gtri
	// Gửi (ch <- 1) sẽ bị chặn khi buffer đầy
	// Nhận (<-ch) sẽ bị chặn khi buffer rỗng
	var wg sync.WaitGroup
	cha := make(chan string, 3)

	for i := range 3 {
		wg.Add(1)
		go task(i, cha, &wg)
	}

	wg.Wait()
	close(cha)

	for val := range cha {
		fmt.Println(val)
	}
}
